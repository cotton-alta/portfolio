<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="card-wrapper" v-for="item in items" :key="item.id">
      <workCard :card="item"/>
    </div>
  </div>
</template>

<script>
import workCard from "~/components/ui/editCard.vue"
import tabTitle from "~/components/layouts/tabTitle.vue"

export default {
  props: ["items"],
  components: {
    workCard,
    tabTitle
  },
  data() {
    let pageTitle = ""
    if(this.$route.path === "/edit/worksList") {
      pageTitle = "作品一覧"
    } else {
      pageTitle = "記事一覧"
    }
    return { pageTitle }
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

.card-wrapper {
  margin-bottom: 30px;
}

</style>