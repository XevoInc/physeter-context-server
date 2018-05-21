import 'isomorphic-fetch';
import moment from 'moment';

const baseUrl =
  'http://physeter-context-server-gateway-569507109.us-west-2.elb.amazonaws.com';

export function fetchRecommends({ lat, lng, datetime, fuel, passenger }) {
  const path = '/api/v1/recommends';
  const params = new URLSearchParams();
  const obj = {
    'car_state.current_location.latitude': lat,
    'car_state.current_location.longitude': lng,
    'car_state.fuel_level_percentage': fuel,
    'car_state.number_of_passengers': passenger,
    time: moment(datetime).format(),
  };
  Object.keys(obj).forEach(key => params.append(key, obj[key]));
  return fetch(baseUrl + path + '?' + params.toString(), { mode: 'cors' }).then(
    res => res.json()
  );
}
