<script setup>
import OrderTable from './components/OrderTable.vue';
</script>

<template>
  <div class="container">
    <OrderTable />
    <br>
  </div>
</template>

<script>
export default {
  components: {
    OrderTable,
  }
}
</script>

<style>
  .container {
    width: 100%;
    border: 1px solid rgb(169, 169, 169);
    border-radius: 20px;
  }
</style>
