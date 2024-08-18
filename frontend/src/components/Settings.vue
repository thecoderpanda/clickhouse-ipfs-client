<template>
    <div>
      <h1>Settings</h1>
      <form @submit.prevent="saveSettings">
        <div>
          <label for="ipfsNodeURL">IPFS Node URL:</label>
          <input v-model="settings.ipfsNodeURL" id="ipfsNodeURL" required />
        </div>
        <div>
          <label for="clickHouseURL">ClickHouse URL:</label>
          <input v-model="settings.clickHouseURL" id="clickHouseURL" required />
        </div>
        <div>
          <label for="clickHouseUser">ClickHouse User:</label>
          <input v-model="settings.clickHouseUser" id="clickHouseUser" required />
        </div>
        <div>
          <label for="clickHousePass">ClickHouse Password:</label>
          <input type="password" v-model="settings.clickHousePass" id="clickHousePass" required />
        </div>
        <button type="submit">Save Settings</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  
  export default {
    data() {
      return {
        settings: {
          ipfsNodeURL: '',
          clickHouseURL: '',
          clickHouseUser: '',
          clickHousePass: ''
        }
      }
    },
    methods: {
      saveSettings() {
        axios.post('http://localhost:8080/settings', this.settings)
          .then(response => {
            alert('Settings saved!')
          })
          .catch(error => {
            console.error(error)
          })
      }
    }
  }
  </script>
  