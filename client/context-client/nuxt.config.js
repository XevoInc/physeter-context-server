const webpack = require('webpack');

module.exports = {
  head: {
    title: 'Physeter Context',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'Nuxt.js project' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },
  loading: { color: '#3B8070' },
  build: {
    vendor: ['iview'],
    extend(config, { isDev, isClient }) {
      config.module.rules.push({
        test: /\.vue$/,
        loader: 'iview-loader',
        options: {
          prefix: false,
        },
      });
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/,
        });
      }

      config.externals = config.externals || {};
      // config.externals['fetch'] = 'fetch';
      config.externals['mapboxgl'] = 'mapboxgl';
    },
  },
  plugins: [{ src: '~/plugins/iview.js', ssr: true }],
  css: ['iview/dist/styles/iview.css'],
  cache: true,
};
