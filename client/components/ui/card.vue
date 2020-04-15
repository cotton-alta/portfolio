<template>
  <div class="card-wrapper">
    <div class="img-holder" :style="`background-image: url(${card.Image})`">
    </div>
    <!-- /.img-holder -->
    <div class="card-text">
      <div v-if="pageJadge" class="card-title">
        <nuxt-link :to="`/articles/${card.Title}`">{{ card.Title }}</nuxt-link>
      </div>
      <!-- /.title -->
      <div v-else class="card-title">
        {{ card.Title }}
      </div>
      <!-- /.title -->
      <div class="card-time">
        {{ modifiedTime }}
      </div>
      <!-- /.card-time -->
      <div class="card-detail">
        {{ card.Detail }}
      </div>
      <!-- /.card-detail -->
      <div v-if="!pageJadge" class="card-link">
        <a :href="card.Link">作品はこちら</a>
      </div>
      <!-- /.card-link -->
      <div v-if="!pageJadge" class="card-github">
        <a :href="card.Github">Git Hub</a>
      </div>
      <!-- /.card-github -->
    </div>
    <!-- /.card-text -->
  </div>
  <!-- /.card-wrapper -->
</template>

<script>
import moment from "moment"

export default {
  props: ["card"],
  data() {
    let pageJadge = false
    if(this.$route.path === "/blog") {
      pageJadge = true
    }
    let modifiedTime = moment(this.card.Timestamp).format('LLL')
    return { 
      pageJadge,
      modifiedTime
    }
  }
}
</script>

<style lang="scss" scoped>

.card {
  &-wrapper {
    color: $string-color;
    margin: 0 auto;
    width: 80%;
    background-color: $black-background;
    border-radius: 5px;
    border: 2px solid $menu-background;
  }
  &-text {
    margin: 10px;
    // padding-bottom: 10px;
  }
  &-title {
    font-size: calc(20px + 0.5vw);
    font-weight: bolder;
    margin-bottom: 10px;
    a {
      color: $string-color;
    }
  }
  &-time {
    margin: 5px 0px;
    line-height: 30px;
  }
  &-detail {
    margin: 5px 0px;
    line-height: 30px;
  }
  &-link {
    color: $string-color;
    font-weight: bolder;
    line-height: 30px;
    a {
      color: $string-color;
    }
  }
  &-github {
    font-weight: bolder;
    line-height: 30px;
    a {
      color: $string-color;
    }
  }
}

.img-holder {
  width: 100%;
  padding-top: 56.25%;
  background-position: center center;
  background-size: cover;
}

</style>