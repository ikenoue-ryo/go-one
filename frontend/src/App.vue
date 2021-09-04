<template>
  <categories />
</template>

<script lang="ts">
import { defineComponent } from "@vue/composition-api";
import categories from './components/categories.vue';

export default defineComponent({
  name: 'App',
  components: {
    categories
  }
});
</script>