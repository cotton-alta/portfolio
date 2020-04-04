<template>
  <div class="btn-wrapper">
    <nuxt-link :to="link.href">{{ link.title }}</nuxt-link>
  </div>
  <!-- /.btn-wrapper -->
</template>

<script>
export default {
  props: ["link"]
}
</script>

<style lang="scss" scoped>
.btn-wrapper {
  width: 200px;
  height: 50px;
  line-height: 50px;
  text-align: center;
  border-radius: 5px;
  background-color: #D0D0D2;
  a {
    display: inline-block;
    width: 100%;
    height: 100%;
    font-size: calc(20px + 0.6vw);
    font-weight: bolder;
    text-decoration: none;
    color: $main-background;
  }
}
</style>