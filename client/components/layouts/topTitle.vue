<template>
  <div class="title-wrapper">
    <span>cotton's portfolio</span>
  </div>
  <!-- /.title-wrapper -->
</template>

<script>
export default {
  
}
</script>

<style lang="scss" scoped>

.title-wrapper {
  height: 100px;
  line-height: 100px;
  text-align: center;
  span {
    font-weight: bolder;
    color: #ffffff;
    font-size: calc(35px + 0.5vw);
    @include mq(xs) {
      font-size: calc(35px + 1.5vw);
    }
  }
}

</style>