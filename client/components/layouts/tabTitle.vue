<template>
  <div class="title-wrapper">
    <span>{{ pageTitle }}</span>
  </div>
  <!-- /.title-wrapper -->
</template>

<script>
export default {
  props: ["pageTitle"]
}
</script>

<style lang="scss" scoped>

.title-wrapper {
  margin: 0 auto;
  padding-top: 10vh;
  margin-bottom: 30px;
  width: 90%;
  text-align: center;
  border-bottom: 3px solid $menu-background;
  padding-bottom: 20px;
  span {
    color: $menu-background;
    font-size: calc(25px + 1.5vw);
    font-weight: bolder;
  }
}

</style>