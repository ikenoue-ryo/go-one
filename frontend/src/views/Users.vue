<template>
  <div class="users">
    <h1>{{ msg }}</h1>
    <!-- {{users}} -->

    <div v-for="user in users" :key="user.id">
      {{user.id}} {{ user.name }} {{ user.email }}
    </div>

    <div>
      <p><input v-model="name" type="text" name="name" placeholder="name"></p>
      <p><input v-model="email" type="mail" name="email" placeholder="email"></p>
      <button type="button" @click="createUser">送信</button>
    </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/composition-api";
import axios from 'axios';

export default defineComponent({
  name: 'Users',
  data(){
    return {
      msg:'Hello users',
      users: '',
      name: '',
      email: '',
  }},
  mounted() {
      axios
        .get('/users')
        .then(response => {
          this.users = response.data
        })
        .catch(err => {
          console.log(err)
        })
  },
  methods: {
    createUser: function(){
      axios.post('/users', {
        name: this.name,
        email: this.email,
      }).then(res => {
        console.log(res);
      }).catch(err => {
        console.log(err);
      })
    }
  }
})
</script>