<template>
  <div id="map" class="map" />
</template>

<script>
import { M_UPDATE_LAT, M_UPDATE_LNG } from '../store/condition';
import { A_FETCH_POIS } from '../store/pois';

const ns = 'condition/';
const poi = 'pois/';

export default {
  head: {
    script: [
      { src: 'https://api.tiles.mapbox.com/mapbox-gl-js/v0.44.2/mapbox-gl.js' }
    ],
    link: [
      { rel: 'stylesheet', href: 'https://api.tiles.mapbox.com/mapbox-gl-js/v0.44.2/mapbox-gl.css' }
    ]
  },
  data() {
    return {
      map: null,
    };
  },
  computed: {
    lat: {
      get() { return this.$store.state.condition.lat; },
      set(value) {
        this.$store.commit(ns + M_UPDATE_LAT, value);
        this.sendRequest();
      },
    },
    lng: {
      get() { return this.$store.state.condition.lng; },
      set(value) {
        this.$store.commit(ns + M_UPDATE_LNG, value);
        this.sendRequest();
      },
    },
  },
  created() {
  },
  mounted() {
    this.initializeMap();
  },
  methods: {
    initializeMap() {
      const mapboxgl = window.mapboxgl;
      mapboxgl.accessToken = 'pk.eyJ1IjoiYW1vcmlubyIsImEiOiJkODQ3MTI2YjBlMTJhZTRmYzIwMDhjYTY1OThiZmExZCJ9.TuMkRWoKBzfRx4nJUx9i8Q';

      this.map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/streets-v9',
        center: [this.lng, this.lat],
        zoom: 14,
      });
    },
    sendRequest() {
      this.$store.dispatch(poi + A_FETCH_POIS);
    },
  },
}
</script>

<style scoped>
.map {
  width: 100%;
  height: 640px;
}
</style>
