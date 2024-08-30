<template>
    <v-container>
      <h2>ClickHouse Data</h2>
      <v-data-table
        :headers="headers"
        :items="items"
        class="elevation-1"
      ></v-data-table>
    </v-container>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data: () => ({
      headers: [
        { text: 'Filename', value: 'filename' },
        { text: 'CID', value: 'cid' },
        { text: 'Data Type', value: 'dataType' },
        { text: 'Data', value: 'data' },
      ],
      items: [],
    }),
    async created() {
      try {
        const response = await axios.get('/api/clickhouse/data');
        this.items = response.data;
      } catch (error) {
        console.error('Failed to fetch ClickHouse data:', error);
      }
    },
  };
  </script>