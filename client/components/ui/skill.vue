<template>
  <div class="skill-wrapper">
    <div class="skill-image">
      <img :src="skill.img" />
    </div>
    <!-- /.skill-image -->
    <p>{{ skill.name }}</p>
    <p>経験年数 : {{ skill.years }}</p>
  </div>
  <!-- /.skill-wrapper -->
</template>

<script>
export default {
  props: ["skill"]
}
</script>

<style lang="scss" scoped>

.skill {
  &-wrapper {
    width: 120px;
    text-align: center;
    font-weight: bolder;
    line-height: 30px;
    color: $menu-background;
  }
  &-image {
    height: 100px;
    line-height: 100px;
    text-align: center;
    img {
    vertical-align: middle;
      width: 80%;
    }
  }
}

</style>