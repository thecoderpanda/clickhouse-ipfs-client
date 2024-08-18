<template>
    <v-app>
      <v-navigation-drawer app>
        <v-list dense>
          <v-list-item-group>
            <v-list-item
              v-for="(item, index) in menuItems"
              :key="index"
              :to="item.route"
              router
            >
              <v-list-item-icon>
                <v-icon>{{ item.icon }}</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>{{ item.title }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-navigation-drawer>
  
      <v-app-bar app color="primary" dark>
        <v-toolbar-title>My Dashboard</v-toolbar-title>
      </v-app-bar>
  
      <v-main>
        <v-container fluid>
          <router-view></router-view>
        </v-container>
      </v-main>
    </v-app>
  </template>
  
  <script>
  export default {
    data: () => ({
      menuItems: [
        { title: 'Dashboard', icon: 'mdi-view-dashboard', route: '/' },
        { title: 'Settings', icon: 'mdi-cog', route: '/settings' },
      ],
    }),
  };
  </script>
  