<template>
    <table>
      <thead>
        <tr>
          <th>Data</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data" :key="index">
          <td>{{ item }}</td>
        </tr>
      </tbody>
    </table>
  </template>
  
  <script>
  export default {
    props: ['data']
  }
  </script>
  