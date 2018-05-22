<template>
  <div>
    <div
      id="map"
      class="map" />
  </div>
</template>

<script>
import mycar from '~/assets/mycar.svg';
import mapPin from '~/assets/map_marker.png';
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
      mapPin,

      map: null,
      markers: [],
      watchCancel: null,
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
        const loc = this.location;
        if (lnglat[0] === loc[0] && lnglat[1] === loc[1]) return;
        this.$store.commit(ns + M_UPDATE_LNG, lnglat[0]);
        this.$store.commit(ns + M_UPDATE_LAT, lnglat[1]);
        this.sendRequest();

        this.map.panTo(lnglat);
      },
    },
  },
  mounted() {
    this.initializeMap();

    this.watchCancel = this.$store.watch(
      () => {
        return this.$store.state.pois.list;
      },
      value => {
        this.updatePins(value);
      }
    );
    this.updatePins(this.$store.state.pois.list);
  },
  beforeDestroy() {
    if (this.watchCancel) {
      this.watchCancel();
    }
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

        const geojson = lnglat => ({
          type: 'Feature',
          geometry: {
            type: 'Point',
            coordinates: lnglat,
          },
        });

        map.addSource(myCarLayerId, {
          type: 'geojson',
          data: geojson(this.location),
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

        const canvas = map.getCanvasContainer();
        const onMove = e => {
          const coords = e.lngLat;
          canvas.style.cursor = 'grabbing';

          const lnglat = [coords.lng, coords.lat];
          map.getSource(myCarLayerId).setData(geojson(lnglat));
        };
        const onUp = e => {
          const coords = e.lngLat;

          this.location = [coords.lng, coords.lat];
          canvas.style.cursor = '';

          map.off('mousemove', onMove);
          map.off('touchmove', onMove);
        };

        map.on('mouseenter', myCarLayerId, () => {
          canvas.style.cursor = 'move';
          map.dragPan.disable();
        });
        map.on('mouseleave', myCarLayerId, () => {
          canvas.style.cursor = '';
          map.dragPan.enable();
        });
        map.on('mousedown', myCarLayerId, e => {
          canvas.style.cursor = 'grab';
          map.on('mousemove', onMove);
          map.once('mouseup', onUp);
        });
        map.on('touchstart', myCarLayerId, e => {
          if (e.points.length !== 1) return;
          map.on('touchmove', onMove);
          map.once('touchend', onUp);
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
    updatePins(list) {
      this.markers.forEach(marker => marker.remove());
      this.markers = list.map((poi, index) => {
        const loc = poi.coordinates.location;
        return this.createPin(
          [loc.longitude, loc.latitude],
          index + 1,
          poi.name
        );
      });
    },
    createPin(lnglat, num, name) {
      const el = document.createElement('div');
      el.className = 'map-marker';
      el.style.backgroundImage = `url(${this.mapPin})`;
      {
        const text = document.createElement('div');
        text.className = 'map-marker-text';
        text.textContent = num;
        el.appendChild(text);
      }
      // el.addEventListener('click', e => {
      //   e.stopPropagation();
      // });

      const popup = new mapboxgl.Popup({ offset: 33 }).setText(name);
      return new mapboxgl.Marker(el)
        .setLngLat(lnglat)
        .setPopup(popup)
        .addTo(this.map);
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

<style>
.map-marker {
  width: 25px;
  height: 41px;
  background-size: 100% 100%;
}
.map-marker-text {
  position: relative;
  width: 25px;
  top: 4px;
  font-size: 12px;
  font-weight: 500;
  text-align: center;
}
</style>
