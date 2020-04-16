<template>
  <div class="container">
    <topTitle />
    <topButton class="btn" :link="link"/>
  </div>
  <!-- /.container -->
</template>

<script>
import topTitle from "~/components/layouts/topTitle.vue"
import topButton from "~/components/ui/button.vue"

export default {
  layout: "toppage",
  components: {
    topTitle,
    topButton
  },
  data() {
    return { link: {title: "CLICK HERE", href: "/profile"} }
  }
}
</script>

<style>
.container {
  min-height: 100vh;
  width: 90%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.btn {
  margin: 10px auto;
}
</style>
