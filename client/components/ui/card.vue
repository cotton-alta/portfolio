<template>
  <div class="card-wrapper">
    <div class="img-holder" :style="`background-image: url(${card.Image})`">
    </div>
    <!-- /.img-holder -->
    <div class="card-text">
      <div class="card-title">
        {{ card.Title }}
      </div>
      <!-- /.title -->
      <div class="card-detail">
        {{ card.Detail }}
      </div>
      <!-- /.card-detail -->
    </div>
    <!-- /.card-text -->
  </div>
  <!-- /.card-wrapper -->
</template>

<script>
export default {
  props: ["card"]
}
</script>

<style lang="scss" scoped>

.card {
  &-wrapper {
    color: $menu-background;
    margin: 0 auto;
    width: 80%;
    background-color: #252526;
    border-radius: 5px;
  }
  &-text {
    height: 100px;
    margin: 10px;
  }
  &-title {
    font-size: calc(20px + 0.5vw);
    font-weight: bolder;
    margin-bottom: 10px;
  }
  &-detail {
    font-weight: bolder;
    line-height: 30px;
  }
}

.img-holder {
  width: 100%;
  padding-top: 56.25%;
  background-position: center center;
  background-size: cover;
  border-radius: 5px;
}

</style>