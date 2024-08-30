
import express from 'express';
import bodyParser from 'body-parser';
import sqlite3 from 'sqlite3';
import { promises as fs } from 'fs';
import multer from 'multer';
import path from 'path';
import { createClient } from '@clickhouse/client'
import { createHelia } from 'helia';
import { unixfs } from '@helia/unixfs';
import { CID } from 'multiformats/cid';

const ClickHouse = createClient({
    url: 'http://localhost:8123',
    basicAuth: {
        username: 'default',
        password: '',
    },
});

const app = express();
const port = 8080;

let helia;
let heliaFs;

(async () => {
    helia = await createHelia();
    heliaFs = unixfs(helia);
})();

// Middleware
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

// Database setup
const db = new sqlite3.Database('./localdb.sqlite', (err) => {
    if (err) {
        console.error('Error opening database', err);
    } else {
        db.run(`CREATE TABLE IF NOT EXISTS settings (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      ipfsNodeUrl TEXT,
      ipfsUsername TEXT,
      ipfsPassword TEXT,
      clickHouseUrl TEXT,
      clickHouseUser TEXT,
      clickHousePass TEXT
    )`);

        db.run(`CREATE TABLE IF NOT EXISTS cid_clickhouse (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      cid TEXT,
      filename TEXT,
      clickhouseUrl TEXT
    )`);
    }
});

// Multer setup for file uploads
const upload = multer({ dest: 'uploads/' });

// Routes
app.get('/api/settings', (req, res) => {
    db.get('SELECT * FROM settings WHERE id = 1', (err, row) => {
        if (err) {
            console.error('Error fetching settings:', err);
            res.status(500).send(err);
        } else {
            res.json(row);
        }
    });
});

app.post('/api/settings', (req, res) => {
    const { ipfsNodeUrl, ipfsUsername, ipfsPassword, clickHouseUrl, clickHouseUser, clickHousePass } = req.body;
    db.run(`INSERT OR REPLACE INTO settings (id, ipfsNodeUrl, ipfsUsername, ipfsPassword, clickHouseUrl, clickHouseUser, clickHousePass) VALUES (1, ?, ?, ?, ?, ?, ?)`,
        [ipfsNodeUrl, ipfsUsername, ipfsPassword, clickHouseUrl, clickHouseUser, clickHousePass],
        (err) => {
            if (err) {
                console.error('Error saving settings:', err);
                res.status(500).send(err);
            } else {
                res.sendStatus(200);
            }
        });
});

app.post('/api/ipfs/upload', upload.single('file'), async (req, res) => {
    const file = req.file;
    if (!file) {
        return res.status(400).send('No file uploaded');
    }

    try {
        const fileData = await fs.readFile(file.path);
        const cid = await heliaFs.addBytes(fileData);

        db.run(`INSERT INTO cid_clickhouse (cid, filename) VALUES (?, ?)`, [cid.toString(), file.originalname], (err) => {
            if (err) {
                console.error('Error inserting CID into database:', err);
                res.status(500).send(err);
            } else {
                res.json({ cid: cid.toString(), filename: file.originalname });
            }
        });
    } catch (error) {
        console.error('Error uploading file to IPFS:', error);
        res.status(500).send(error);
    } finally {
        await fs.unlink(file.path);
    }
});

app.post('/api/ipfs/fetch', async (req, res) => {
    const { cid } = req.body;
    if (!cid) {
        return res.status(400).send('CID is required');
    }

    try {
        const parsedCid = CID.parse(cid);
        const chunks = [];
        for await (const chunk of heliaFs.cat(parsedCid)) {
            chunks.push(chunk);
        }
        const data = Buffer.concat(chunks).toString();
        res.json({ cid, data });
    } catch (error) {
        console.error('Error fetching data from IPFS:', error);
        res.status(500).send(error);
    }
});

app.post('/api/clickhouse/upload', (req, res) => {
    const { cid, filename, data } = req.body;
    db.get('SELECT * FROM settings WHERE id = 1', async (err, settings) => {
        if (err) {
            console.error('Error fetching settings:', err);
            res.status(500).send(err);
        } else {
            try {
                const clickhouse = new ClickHouse({
                    url: settings.clickHouseUrl,
                    basicAuth: {
                        username: settings.clickHouseUser,
                        password: settings.clickHousePass,
                    },
                });
                await clickhouse.query(`INSERT INTO ipfs_data (cid, filename, data) VALUES ('${cid}', '${filename}', '${data}')`).toPromise();
                db.run(`UPDATE cid_clickhouse SET clickhouseUrl = ? WHERE cid = ?`, [settings.clickHouseUrl, cid], (err) => {
                    if (err) {
                        console.error('Error updating CID in database:', err);
                        res.status(500).send(err);
                    } else {
                        res.sendStatus(200);
                    }
                });
            } catch (error) {
                console.error('Error uploading data to ClickHouse:', error);
                res.status(500).send(error);
            }
        }
    });
});

app.get('/api/clickhouse/data', (req, res) => {
    db.get('SELECT * FROM settings WHERE id = 1', async (err, settings) => {
        if (err) {
            console.error('Error fetching settings:', err);
            res.status(500).send(err);
        } else {
            try {
                const clickhouse = new ClickHouse({
                    url: settings.clickHouseUrl,
                    basicAuth: {
                        username: settings.clickHouseUser,
                        password: settings.clickHousePass,
                    },
                });
                const rows = await clickhouse.query('SELECT * FROM ipfs_data').toPromise();
                res.json(rows);
            } catch (error) {
                console.error('Error fetching data from ClickHouse:', error);
                res.status(500).send(error);
            }
        }
    });
});

// Error handling middleware
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({
        message: 'An unexpected error occurred',
        error: process.env.NODE_ENV === 'production' ? {} : err
    });
});

app.listen(port, () => {
    console.log(`Server running on http://localhost:${port}`);
});