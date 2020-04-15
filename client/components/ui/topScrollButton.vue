<template>
  <transition name="scroll">
    <div class="scroll-wrapper">
      <a v-scroll-to="'#top'"></a>
    </div>
    <!-- /.scroll-wrapper -->
  </transition>
</template>

<script>
export default {
 transition: {
   name: 'scroll'
 }
}
</script>

<style lang="scss" scoped>

.scroll-wrapper {
  cursor: pointer;
  width: 3rem;
  height: 3rem;
  opacity: 0.6;
  background: $menu-background;
  border-radius: 50%;
  a {
    display: block;
    position: relative;
    height: 100%;
    width: 100%;
  }
  a::before {
    position: absolute;
    top: 0.25rem;
    left: 0;
    right: 0;
    bottom: 0;
    width: 100%;
    height: 100%;
    font-size: 40px;
    font-weight: 900;
    font-family: 'Font Awesome 5 Free';
    content: '\f106';
    text-align: center;
    color: $main-background;
  }
}

.scroll-enter-active, .scroll-leave-active {
    transition: opacity .5s;
}
.scroll-enter, .scroll-leave-to {
    opacity: 0;
}

</style>