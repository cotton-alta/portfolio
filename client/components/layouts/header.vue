<template>
  <div class="header-wrapper">
    <div v-for="data in datas" :key="data.id" class="header-card">
      <nuxt-link class="header-name" :to="data.href">{{ data.title }}</nuxt-link>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      datas: [
        { title: "TOP", href: "/" },
        { title: "PROFILE", href: "/profile" },
        { title: "SKILLS", href: "/skills" },
        { title: "WORKS", href: "/works" },
        { title: "BLOG", href: "/blog" },
        { title: "CONTACT", href: "/contact" },
      ]
    }
  }
}
</script>

<style lang="scss" scoped>

.header {
  &-wrapper {
    display: none;
    @include mq {
      height: 80px;
      line-height: 80px;
      width: 100%;
      background-color: $menu-background;
      display: flex;
      flex-direction: row;
      justify-content: center;
      justify-items: center;
    }
  }
  &-card {
    display: inline-block;
    margin: 0 15px;
    vertical-align: middle;
  }
  &-name {
    text-decoration: none;
    font-size: calc(20px + 0.8vw);
    font-weight: bolder;
    color: $main-background;
    &:hover{
      color: #ffffff;
    }
  }
}

</style>