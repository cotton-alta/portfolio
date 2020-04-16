<template>
  <div>
    <div class="header-wrapper">
      <div class="header-card">
        <!-- <span class="header-name" @click="linkToTop">TOP</span> -->
        <a href="/" class="header-name">TOP</a>
      </div>
      <div v-for="data in datas" :key="data.id" class="header-card">
        <nuxt-link class="header-name" :to="data.href">{{ data.title }}</nuxt-link>
      </div>
      <!-- /.header-card -->
    </div>
    <!-- /.header-wrapper -->
    <div class="hamburger-wrapper">
      <div 
        :class="{
          'hamburger-icon--close': !isOpen,
          'hamburger-icon--open': isOpen
        }"
        @click="hamburgerOpen"
      >
        <p></p>
        <span></span>
      </div>
      <!-- /.hamburger-icon -->
      <div 
        :class="{
          'hamburger-content--open': isOpen,
          'hamburger-content--close': !isOpen
        }"
      >
        <div class="hamburger-flex">
          <div class="hamburger-card">
            <a class="hamaburger-name" href="/">TOP</a>
          </div>
          <div
            class="hamburger-card"
            v-for="data in datas"
            :key="data.id"
            @click="hamburgerOpen"
          >
            <nuxt-link class="hamaburger-name" :to="data.href">
              {{ data.title }}
            </nuxt-link>
          </div>
          <!-- /.hamburger-card -->
        </div>
        <!-- /.hamburger-flex -->
      </div>
      <!-- /.hamburger-content -->
    </div>
    <!-- /.hamburger-wrapper -->
  </div>
</template>

<script>
export default {
  data() {
    return {
      datas: [
        { title: "PROFILE", href: "/profile" },
        { title: "SKILLS", href: "/skills" },
        { title: "WORKS", href: "/works" },
        { title: "BLOG", href: "/blog" },
        { title: "CONTACT", href: "/contact" },
      ],
      isOpen: false
    }
  },
  mounted() {
    window.addEventListener("scroll", () => {
      this.isOpen = false
    })
  },
  methods: {
    hamburgerOpen: function() {
      this.isOpen = !this.isOpen
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

.hamburger {
  &-wrapper {
    height: 50px;
    line-height: 60px;
    width: 100%;
    padding-right: 16px;
    text-align: right;
    @include mq {
      display: none;
    }
  }
  &-icon {
    transition: all 0.3s;
  }
  &-icon--close {
    display: inline-block;
    width: 32px;
    height: 34px;
    z-index: 9999;
    vertical-align: middle;
    cursor: pointer;
    span, span::before, span::after {
      position: absolute;
      height: 4px;
      width: 32px;
      background: $menu-background;
      display: block;
      border-radius: 3px;
      content: "";
    }
    span::before {
      bottom: -13px;
    }
    span::after {
      bottom: -26px;
    }
  }
  &-icon--open {
    display: inline-block;
    width: 32px;
    height: 34px;
    vertical-align: middle;
    cursor: pointer;
    span, span::before, span::after {
      transition: all 0.3s;
      z-index: 9999;
      position: absolute;
      height: 4px;
      width: 32px;
      background: $menu-background;
      display: block;
      border-radius: 3px;
      content: "";
    }
    span::before {
      transform: rotate(45deg);
      background: $main-background;
      bottom: -11px;
    }
    span::after {
      bottom: -11px;
      background: $main-background;
      transform: rotate(-45deg);
    }
    p {
      position: absolute;
      z-index: 1000;
      width: 32px;
      height: 34px;
      content: "";
      cursor: pointer;
    }
  }
  &-content {
    &--open, &--close {
      position: fixed;
      background: $menu-background;
      top: 0;
      left: 20%;
      height: 100%;
      width: 80%;
    }
    &--open {
      transform: translateX(0%);
      transition: .3s ease-in-out;
    }
    &--close {
      transform: translateX(100%);
      transition: .3s ease-in-out;
    }
  }
  &-flex {
    height: 60%;
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    margin: auto;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
  }
  &-card {
    text-align: center;
    a {
      font-size: calc(30px + 1vw);
      font-weight: bolder;
      color: $main-background;
      text-decoration: none;
    }
  }
}

</style>