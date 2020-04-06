<template>
  <div class="card-wrapper">
    <div class="img-holder" :style="`background-image: url(${profileImage})`">
    </div>
    <!-- /.img-holder -->
    <div class="card-text">
      <div class="card-title">
        タイトル
      </div>
      <!-- /.title -->
      <div class="card-detail">
        説明です。
      </div>
      <!-- /.card-detail -->
    </div>
    <!-- /.card-text -->
  </div>
  <!-- /.card-wrapper -->
</template>

<script>
import profileImage from "~/static/profileImage.jpg"

export default {
  data() {
    return { profileImage }
  }
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
    font-size: calc(15px + 0.4vw);
    font-weight: bolder;
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