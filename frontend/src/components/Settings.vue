<template>
  <v-container>
    <v-card>
      <v-card-title>
        <span class="headline">Settings</span>
      </v-card-title>
      <v-card-text>
        <v-form>
          <v-text-field label="IPFS Node URL" v-model="settings.ipfsNodeUrl"></v-text-field>
          <v-text-field label="IPFS Username" v-model="settings.ipfsUsername"></v-text-field>
          <v-text-field label="IPFS Password" v-model="settings.ipfsPassword" type="password"></v-text-field>
          <v-text-field label="ClickHouse URL" v-model="settings.clickHouseUrl"></v-text-field>
          <v-text-field label="ClickHouse User" v-model="settings.clickHouseUser"></v-text-field>
          <v-text-field label="ClickHouse Password" v-model="settings.clickHousePass" type="password"></v-text-field>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="primary" @click="saveSettings">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  data: () => ({
    settings: {
      ipfsNodeUrl: '',
      ipfsUsername: '',
      ipfsPassword: '',
      clickHouseUrl: '',
      clickHouseUser: '',
      clickHousePass: '',
    },
  }),
  methods: {
    async saveSettings() {
      try {
        await axios.post('/api/settings', this.settings);
        alert('Settings saved successfully');
      } catch (error) {
        console.error('Save failed:', error);
      }
    },
  },
  async created() {
    try {
      const response = await axios.get('/api/settings');
      this.settings = response.data;
    } catch (error) {
      console.error('Failed to load settings:', error);
    }
  },
};
</script>