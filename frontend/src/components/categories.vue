<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    {{categories}}

    <div v-for="category in categories" :key="category.id">
      {{category.id}} {{ category.name }} {{ category.kana }}
    </div>


  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/composition-api";
import axios from 'axios';

export default defineComponent({
  name: 'HelloWorld',
  data(){
    return {
      msg:'Hello TypeScript',
      message: '',
      categories: ''
  }},
  mounted() {
      axios
        .get('/category')
        .then(response => {
          this.categories = response.data
        })
        .catch(err => {
          console.log(err)
        })
  }
})
</script>