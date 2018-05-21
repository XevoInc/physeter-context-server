<template>
  <div class="container">
    <h2>
      <Icon type="star" />
      Recommendation List
    </h2>

    <Table 
      :loading="loading" 
      :columns="columns" 
      :data="pois" 
      highlight-row>
      <div slot="loading">
        <Spin>
          <Icon 
            type="load-c" 
            size="18" 
            class="spin-icon-load"/>
          <div>Loading</div>
        </Spin>
      </div>
    </Table>

  </div>
</template>

<script>
export default {
  data() {
    return {
      columns: [
        {
          type: 'index',
          width: 60,
          align: 'center',
        },
        {
          title: 'Name',
          key: 'name',
        },
        {
          title: 'Distance (m)',
          width: 120,
          key: 'distance',
        },
      ],
    };
  },
  computed: {
    pois() {
      return this.$store.state.pois.list;
    },
    loading() {
      return this.$store.state.pois.loading;
    },
  },
};
</script>

<style scoped>
.container {
  margin: 5%;
  margin-top: 0;
}

.spin-icon-load {
  animation: animation-spin 1s linear infinite;
}
@keyframes animation-spin {
  from {
    transform: rotate(0deg);
  }
  50% {
    transform: rotate(180deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
