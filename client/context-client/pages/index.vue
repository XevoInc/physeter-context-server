<template>
  <section class="frame">
    <section class="contents">
      <mapbox/>
    </section>
    <section class="sidebar">
      <condition-panel />
      <result-list />
    </section>
  </section>
</template>

<script>
import ConditionPanel from '~/components/ConditionPanel.vue';
import Mapbox from '~/components/Mapbox.vue';
import ResultList from '~/components/ResultList.vue';

export default {
  components: {
    ConditionPanel,
    Mapbox,
    ResultList,
  }
};
</script>

<style>
.frame {
  display: flex;
}

.contents {
  width: 60%;
}

.sidebar {
  width: 40%;
  display: flex;
  flex-direction: column;
}
</style>
