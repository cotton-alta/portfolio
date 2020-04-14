<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
  </div>
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"

export default {
  components: {
    tabTitle
  },
  data() {
    return { pageTitle: "作品編集" }
  }
}
</script>

<style lang="scss" scoped>
.container {
  display: block;
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
  min-height: calc(100vh - 50px);
  @include mq {
    min-height: calc(100vh - 80px);
  }
}
</style>