<template>
  <div>
    <div
      id="map"
      class="map" />
  </div>
</template>

<script>
import mycar from '~/assets/mycar.svg';
const myCarImageId = 'MyCarImage';
const myCarLayerId = 'myCar';

import { M_UPDATE_LAT, M_UPDATE_LNG } from '~/store/condition';
import { A_FETCH_POIS } from '~/store/pois';
const ns = 'condition/';
const poi = 'pois/';

export default {
  head: {
    script: [
      { src: 'https://api.tiles.mapbox.com/mapbox-gl-js/v0.44.2/mapbox-gl.js' },
    ],
    link: [
      {
        rel: 'stylesheet',
        href: 'https://api.tiles.mapbox.com/mapbox-gl-js/v0.44.2/mapbox-gl.css',
      },
    ],
  },
  data() {
    return {
      mycar,
      map: null,
    };
  },
  computed: {
    location: {
      get() {
        return [
          this.$store.state.condition.lng,
          this.$store.state.condition.lat,
        ];
      },
      set(lnglat) {
        this.$store.commit(ns + M_UPDATE_LNG, lnglat[0]);
        this.$store.commit(ns + M_UPDATE_LAT, lnglat[1]);
        this.sendRequest();
      },
    },
  },
  created() {},
  mounted() {
    this.initializeMap();
  },
  methods: {
    initializeMap() {
      const mapboxgl = window.mapboxgl;
      mapboxgl.accessToken =
        'pk.eyJ1IjoiYW1vcmlubyIsImEiOiJkODQ3MTI2YjBlMTJhZTRmYzIwMDhjYTY1OThiZmExZCJ9.TuMkRWoKBzfRx4nJUx9i8Q';

      this.map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/streets-v9',
        center: this.location,
        zoom: 15,
      });

      this.map.addControl(new mapboxgl.NavigationControl(), 'top-left');
      this.map.addControl(new mapboxgl.ScaleControl(), 'bottom-right');

      const geoLocate = new mapboxgl.GeolocateControl({
        positionOptions: {
          enableHighAccuracy: true,
        },
        trackUserLocation: true,
        showUserLocation: false,
      });
      geoLocate.on('geolocate', event => {
        this.location = [event.coords.longitude, event.coords.latitude];
      });
      this.map.addControl(geoLocate);

      this.map.on('load', () => {
        this.map.resize();
        this.addMyCarLayer();
      });
    },
    addMyCarLayer() {
      return this.loadMyCarImage().then(() => {
        const map = this.map;

        if (map.getLayer(myCarLayerId)) {
          map.removeSource(myCarLayerId);
          map.removeLayer(myCarLayerId);
        }

        map.addSource(myCarLayerId, {
          type: 'geojson',
          data: {
            type: 'Feature',
            geometry: {
              type: 'Point',
              coordinates: this.location,
            },
            properties: null,
          },
        });
        map.addLayer({
          id: myCarLayerId,
          type: 'symbol',
          source: myCarLayerId,
          layout: {
            'icon-image': myCarImageId,
            'icon-rotation-alignment': 'map',
            'icon-size': 0.4,
          },
        });
      });
    },
    loadMyCarImage() {
      const map = this.map;
      if (map.hasImage(myCarImageId)) {
        map.removeImage(myCarImageId);
      }

      return new Promise((resolve, _reject) => {
        const img = document.createElement('img');
        img.onload = () => {
          map.addImage(myCarImageId, img);
          resolve();
        };
        img.src = this.mycar;
      });
    },

    sendRequest() {
      this.$store.dispatch(poi + A_FETCH_POIS);
    },
  },
};
</script>

<style scoped>
.map {
  width: 100%;
  height: 800px;
}
</style>
