<template>
    <v-container>
      <v-card>
        <v-card-title>
          <span class="headline">Settings</span>
        </v-card-title>
        <v-card-text>
          <v-form>
            <v-text-field label="IPFS Node URL" v-model="settings.ipfsNodeUrl"></v-text-field>
            <v-text-field label="ClickHouse URL" v-model="settings.clickHouseUrl"></v-text-field>
            <!-- Add other settings fields as needed -->
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
  export default {
    data: () => ({
      settings: {
        ipfsNodeUrl: '',
        clickHouseUrl: '',
      },
    }),
    methods: {
      async saveSettings() {
        // Save settings to backend
        try {
          await fetch('/api/settings', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(this.settings),
          });
          alert('Settings saved successfully');
        } catch (error) {
          console.error('Save failed:', error);
        }
      },
    },
    async created() {
      // Load current settings from backend
      try {
        const response = await fetch('/api/settings');
        this.settings = await response.json();
      } catch (error) {
        console.error('Failed to load settings:', error);
      }
    },
  };
  </script>
  