<script setup>
import { ref, reactive } from 'vue'
import { DatePicker } from 'v-calendar'
import 'v-calendar/dist/style.css'

const date = ref(new Date())
date.value.setDate(Number(date.value.getDate()))
const range = reactive({
  start: new Date(),
  end: date.value
})
</script>

<template>
    <br>
<div class="topbar">
    <a href="#">
        <img class="searchicon" src = './icons/search.png' style='vertical-align: middle' />&nbsp;&nbsp;Search
    </a>
    <input type="text" v-model="searchValue" placeholder="Search..."/>
</div>
<br>

<div class="datebar">
    <p>Created Date</p>
      <DatePicker v-model="range" mode="date" :columns="2" is-range>
      <template v-slot="{ inputValue, inputEvents }">
        <div class="group">
          <label class="button groupItem" for="start">Start</label>
          <input type="text" id="start" onfocus="this.value=''" :value="inputValue.start" v-on="inputEvents.start" class="input groupItem" :v-model="startDate" placeholder="Start">
          <label class="button groupItem" for="end"> End</label>
          <input type="text" id="end" :value="inputValue.end" readonly class="input groupItem" :v-model="endDate" placeholder="End">
        </div>
      </template>
    </DatePicker>
</div>
<br>
<div class="table-responsive">
    <table>
        <tr>
            <th>Total Amount: ${{ totalRequest }}</th>
        </tr>
        <br>
        <tr>
            <th>Order Name</th>
            <th>Customer Company</th>
            <th>Customer Name</th>
            <th>Ordered Date</th>
            <th>Delivered Amount</th>
            <th>Total Amount</th>
        </tr>
        <tr v-for="order in ordersList" :key="order" v-if="ordersList.length">
            <td>{{ order.order_name }}</td>
            <td>{{ order.customer_company }}</td>
            <td>{{ order.customer_name }}</td>
            <td>{{ order.order_date }}</td>
            <td>${{ order.delivered_amount }}</td>
            <td>${{ order.total_amount }}</td>
        </tr>
        <tr v-else class="alert"><td>No Data Found!</td></tr>
    </table>
    <br>
  </div>
</template>

<script>
import "bootstrap/dist/css/bootstrap.min.css";

export default {    
    name: 'OrderTable',
    data: () => ({
        startDate: null,
        endDate: null,
        searchValue: '',
        orders: [],
        orders_total_amount: 0
  }),
  methods: {
        async  getData() {
            const res = await fetch("http://127.0.0.1:3000/api/orders", {
                        method: 'GET',
                        headers: {
                          'Content-Type': 'application/json'
                        }
                      });
            const finalRes = await res.json();
            this.orders = finalRes;
        }
    },
    mounted() {
        this.getData()
    },
    computed: {
        ordersList() {
            if (this.searchValue.trim().length > 0) {
                // for(i in this.orders) { this.orders_total_amount += this.orders[i].total_amount; }
                return this.orders.filter((order) => order.order_name.toLowerCase().includes(this.searchValue.trim().toLowerCase()))
            }
            // for(i in this.orders) { this.orders_total_amount += this.orders[i].total_amount; }
            return this.orders
        },
        totalRequest() {
          return this.orders.reduce( (acc, item) => {
                return acc + item.total_amount
              }, 0)
        }
    }
  }
</script>

<style>

.searchicon {
        width: 30px;
    }

    
.topbar {
  overflow: hidden;
}
 p {
    font-size: 17px;
 }

.topbar a {
  float: left;
  display: block;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 25px;
  font-weight: bolder;
}

.topbar a:hover {
    color: #2c3e50;
}


.topbar input[type=text] {
  width: 85%;
  padding: 6px;
  margin-top: 12px;
  margin-right: 16px;
  font-size: 17px;
  border: 1px solid #ccc; 
  border-radius: 5px;;
}

.datebar {
  display: block;
  padding: 14px 16px;
  text-decoration: none;
  font-weight: bolder;
    }
.alert td {
    color: rgb(162, 36, 36);
    border-bottom: none;
    font-weight: bolder;
}

@media screen and (max-width: 600px) {
  .topbar a, .topbar input[type=text] {
    float: none;
    display: block;
    text-align: left;
    width: 100%;
    margin: 0;
    padding: 14px;
  }

  .topbar a {
    display: none;
  }
  
  .topbar input[type=text] {
    border: 1px solid #ccc;  
  }


}

table {
  border-collapse: collapse;
  width: 100%;

  text-decoration: none;
  font-weight: bolder;
  
}
tr:nth-child(even) {
  background-color: #f2f2f2;
}

th {
  padding: 15px;
  text-align: left;
  color: rgb(65, 64, 64);
  font-weight: bold;
  font-size: 17.5px;
}

td {
    padding: 15px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.table-responsive {
    overflow-x:auto;
}

@use form {
  field: label, group, input, button;
}

</style>