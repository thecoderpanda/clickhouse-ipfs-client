<template>
    <div>
      <h1>Dashboard</h1>
      <div>
        <h2>Upload to IPFS</h2>
        <input type="file" @change="uploadFile"/>
      </div>
      <div>
        <h2>Fetch from IPFS</h2>
        <input v-model="cid" placeholder="Enter CID"/>
        <button @click="fetchData">Fetch and Ingest</button>
      </div>
      <div>
        <h2>Data from ClickHouse</h2>
        <DataTable :data="clickhouseData"/>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  import DataTable from './DataTable.vue'
  
  export default {
    components: { DataTable },
    data() {
      return {
        cid: '',
        clickhouseData: []
      }
    },
    methods: {
      uploadFile(event) {
        const file = event.target.files[0]
        const formData = new FormData()
        formData.append('file', file)
  
        axios.post('http://localhost:8080/ipfs/upload', formData)
          .then(response => {
            alert('File uploaded to IPFS. CID: ' + response.data.cid)
          })
          .catch(error => {
            console.error(error)
          })
      },
      fetchData() {
        axios.get(`http://localhost:8080/ipfs/fetch/${this.cid}`)
          .then(response => {
            return axios.post('http://localhost:8080/clickhouse/ingest', {
              cid: this.cid,
              data: response.data.data
            })
          })
          .then(() => {
            return axios.get(`http://localhost:8080/clickhouse/query?cid=${this.cid}`)
          })
          .then(response => {
            this.clickhouseData = response.data.data
          })
          .catch(error => {
            console.error(error)
          })
      }
    }
  }
  </script>
  