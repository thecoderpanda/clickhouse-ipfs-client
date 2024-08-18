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
          ></v-data-table>
        </v-col>
      </v-row>
  
      <!-- Upload Dialog -->
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
  
      <!-- Fetch Dialog -->
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
    </v-container>
  </template>
  
  <script>
  export default {
    data: () => ({
      showUploadDialog: false,
      showFetchDialog: false,
      cid: '',
      selectedFile: null,
      headers: [
        { text: 'Name', value: 'name' },
        { text: 'CID', value: 'cid' },
        { text: 'Size', value: 'size' },
      ],
      items: [],
    }),
    methods: {
      onFileSelected(file) {
        this.selectedFile = file;
      },
      async uploadFile() {
        if (!this.selectedFile) return;
  
        // Handle file upload to IPFS
        const formData = new FormData();
        formData.append('file', this.selectedFile);
  
        try {
          const response = await fetch('/api/ipfs/create', {
            method: 'POST',
            body: formData,
          });
          const result = await response.json();
          this.items.push({
            name: this.selectedFile.name,
            cid: result.cid,
            size: this.selectedFile.size,
          });
          this.showUploadDialog = false;
        } catch (error) {
          console.error('Upload failed:', error);
        }
      },
      async fetchData() {
        if (!this.cid) return;
  
        // Fetch data from IPFS by CID
        try {
          const response = await fetch('/api/ipfs/fetch', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ cid: this.cid }),
          });
          const result = await response.json();
          this.items.push({
            name: result.name,
            cid: this.cid,
            size: result.size,
          });
          this.showFetchDialog = false;
        } catch (error) {
          console.error('Fetch failed:', error);
        }
      },
    },
  };
  </script>
  