<template>
  <div class="card-wrapper">
    <div class="img-holder" :style="`background-image: url(${profileImage})`">
    </div>
    <!-- /.img-holder -->
    <div class="card-text">
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
    margin: 0 auto;
    width: 80%;
    background-color: #252526;
    border-radius: 5px;
  }
  &-text {
    height: 100px;
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