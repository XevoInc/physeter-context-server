export const M_UPDATE_LAT = 'updateLat';
export const M_UPDATE_LNG = 'updateLng';
export const M_UPDATE_DATETIME = 'updateDatetime';
export const M_UPDATE_FUEL = 'updateFuel';
export const M_UPDATE_PASSENGER = 'updatePassenger';

export const state = () => ({
  lat: 35.6505,
  lng: 139.7099,
  datetime: new Date(),
  fuel: 50,
  passenger: 1,
});

export const mutations = {
  [M_UPDATE_LAT](state, value) {
    state.lat = value;
  },
  [M_UPDATE_LNG](state, value) {
    state.lng = value;
  },
  [M_UPDATE_DATETIME](state, value) {
    console.log('UPDATE_DATETIME', value);
    state.datetime = value;
  },
  [M_UPDATE_FUEL](state, value) {
    state.fuel = value;
  },
  [M_UPDATE_PASSENGER](state, value) {
    state.passenger = value;
  },
};
