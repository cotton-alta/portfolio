<template>
  <div class="header-wrapper">
    <div v-for="data in datas" :key="data.id" class="header-card">
      <nuxt-link :to="data.href">{{ data.title }}</nuxt-link>
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
        { title: "SKILLS", href: "/" },
        { title: "WORKS", href: "/" },
        { title: "BLOG", href: "/" },
        { title: "CONTACT", href: "/" },
      ]
    }
  }
}
</script>

<style lang="scss" scoped>

.header-wrapper {
  &-wrapper {
    height: 80px;
    width: 100%;
    background-color: $menu-background;
    display: flex;
    flex-direction: row;
    justify-content: center;
  }
  &-card {
    height: 100%;
  }

}

</style>