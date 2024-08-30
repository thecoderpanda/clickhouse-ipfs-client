<template>
  <v-container>
    <v-row>
      <v-col>
        <v-btn color="primary" @click="showUploadDialog = true">
          Upload Data to IPFS
        </v-btn>
        <v-btn color="secondary" @click="showFetchDialog = true">
          Fetch Data by CID
        </v-btn>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <v-data-table
          :headers="headers"
          :items="items"
          class="elevation-1"
        >
          <template v-slot:item.actions="{ item }">
            <v-btn small color="primary" @click="uploadToClickHouse(item)">
              Upload to ClickHouse
            </v-btn>
          </template>
        </v-data-table>
      </v-col>
    </v-row>

    <v-dialog v-model="showUploadDialog" max-width="500px">
      <v-card>
        <v-card-title>
          <span class="headline">Upload Data to IPFS</span>
        </v-card-title>
        <v-card-text>
          <v-file-input
            label="Select File"
            @change="onFileSelected"
            outlined
            dense
          ></v-file-input>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="uploadFile">Upload</v-btn>
          <v-btn text @click="showUploadDialog = false">Cancel</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="showFetchDialog" max-width="500px">
      <v-card>
        <v-card-title>
          <span class="headline">Fetch Data by CID</span>
        </v-card-title>
        <v-card-text>
          <v-text-field
            label="Enter CID"
            v-model="cid"
            outlined
            dense
          ></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="fetchData">Fetch</v-btn>
          <v-btn text @click="showFetchDialog = false">Cancel</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-btn color="info" @click="viewClickHouseData" class="mt-4">
      View ClickHouse Data
    </v-btn>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  data: () => ({
    showUploadDialog: false,
    showFetchDialog: false,
    cid: '',
    selectedFile: null,
    headers: [
      { text: 'Filename', value: 'filename' },
      { text: 'CID', value: 'cid' },
      { text: 'Data Type', value: 'dataType' },
      { text: 'Actions', value: 'actions', sortable: false },
    ],
    items: [],
  }),
  methods: {
    onFileSelected(file) {
      this.selectedFile = file;
    },
    async uploadFile() {
      if (!this.selectedFile) return;

      const formData = new FormData();
      formData.append('file', this.selectedFile);

      try {
        const response = await axios.post('/api/ipfs/upload', formData);
        this.items.push(response.data);
        this.showUploadDialog = false;
      } catch (error) {
        console.error('Upload failed:', error);
      }
    },
    async fetchData() {
      if (!this.cid) return;

      try {
        const response = await axios.post('/api/ipfs/fetch', { cid: this.cid });
        this.items.push(response.data);
        this.showFetchDialog = false;
      } catch (error) {
        console.error('Fetch failed:', error);
      }
    },
    async uploadToClickHouse(item) {
      try {
        await axios.post('/api/clickhouse/upload', item);
        alert('Data uploaded to ClickHouse successfully');
      } catch (error) {
        console.error('ClickHouse upload failed:', error);
      }
    },
    viewClickHouseData() {
      this.$router.push('/clickhouse-data');
    },
  },
};
</script>
  