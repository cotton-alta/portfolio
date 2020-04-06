<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <workCard />
    <workCard />
    <workCard />
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import workCard from "~/components/ui/card.vue"

export default {
  components: {
    tabTitle,
    workCard
  },
  data() {
    return { pageTitle: "WORKS" }
  }
}
</script>

<style lang="scss" scoped>

.container {
  display: block;
  min-height: calc(100vh - 80px);
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
}
</style>