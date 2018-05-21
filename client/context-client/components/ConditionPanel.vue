<template>
  <div class="container">
    <Card
      title="Condition"
      icon="search">

      <Form :label-width="80">
        <FormItem label="Date and time">
          <DatePicker
            v-model="date"
            type="date"
            placeholder="Select date" />
          <TimePicker
            :steps="[1, 5]"
            v-model="time"
            format="HH:mm"
            placeholder="Select time" />
        </FormItem>
        <FormItem label="Fuel">
          <Slider
            v-model.number="fuel"
            :step="5"
            show-input />
        </FormItem>
        <FormItem label="Passengers">
          <Slider
            v-model.number="passenger"
            :min="1"
            :max="6"
            show-input
            show-stops />
        </FormItem>
      </Form>

    </Card>
  </div>
</template>

<script>
import moment from 'moment';
import {
  M_UPDATE_DATETIME,
  M_UPDATE_FUEL,
  M_UPDATE_PASSENGER,
} from '~/store/condition';
import { A_FETCH_POIS } from '~/store/pois';

const ns = 'condition/';
const poi = 'pois/';

export default {
  data() {
    return {};
  },
  computed: {
    date: {
      get() {
        return this.$store.state.condition.datetime;
      },
      set(value) {
        this.$store.commit(
          ns + M_UPDATE_DATETIME,
          this.mergeDateTime(value, this.time)
        );
        this.sendRequest();
      },
    },
    time: {
      get() {
        const m = moment(this.$store.state.condition.datetime);
        return m.format('HH:mm');
      },
      set(value) {
        this.$store.commit(
          ns + M_UPDATE_DATETIME,
          this.mergeDateTime(this.date, value)
        );
        this.sendRequest();
      },
    },
    fuel: {
      get() {
        return this.$store.state.condition.fuel;
      },
      set(value) {
        this.$store.commit(ns + M_UPDATE_FUEL, value);
        this.sendRequest();
      },
    },
    passenger: {
      get() {
        return this.$store.state.condition.passenger;
      },
      set(value) {
        this.$store.commit(ns + M_UPDATE_PASSENGER, value);
        this.sendRequest();
      },
    },
  },
  created() {
    this.sendRequest();
  },
  methods: {
    mergeDateTime(date, time) {
      if (!date || !time) {
        return new Date();
      }
      const m = moment(date);
      const match = time.match(/^(\d{2})\:(\d{2})$/);
      if (match) {
        m.hours(match[1]).minutes(match[2]);
      }
      return m.toDate();
    },
    sendRequest() {
      this.$store.dispatch(poi + A_FETCH_POIS);
    },
  },
};
</script>

<style scoped>
.container {
  margin: 5%;
}
</style>
