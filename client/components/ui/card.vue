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
export default {
  props: ["card"],
  data() {
    let pageJadge = false
    if(this.$route.path === "/blog") {
      pageJadge = true
    }
    return { pageJadge }
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
    border: 3px solid $menu-background;
  }
  &-text {
    margin: 10px;
    padding-bottom: 15px;
  }
  &-title {
    font-size: calc(20px + 0.5vw);
    font-weight: bolder;
    margin-bottom: 10px;
    a {
      color: $menu-background;
    }
  }
  &-detail {
    font-weight: bolder;
    line-height: 30px;
  }
  &-link {
    color: $menu-background;
    font-weight: bolder;
    line-height: 30px;
    a {
      color: $menu-background;
    }
  }
  &-github {
    font-weight: bolder;
    line-height: 30px;
    a {
      color: $menu-background;
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