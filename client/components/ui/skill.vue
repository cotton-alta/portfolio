<template>
  <div class="skill-wrapper">
    <div class="skill-image">
      <img :src="skill.img" />
    </div>
    <!-- /.skill-image -->
    <span>{{ skill.name }}</span>
    <p>経験年数 : {{ skill.years }}</p>
    <p v-html="skill.level"></p>
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
    color: $string-color;
    span {
      line-height: 33px;
      font-weight: bolder;
      font-size: calc(17px + 0.3vw);
    }
    p {
      line-height: 23px;
    }
  }
  &-image {
    height: 100px;
    line-height: 100px;
    text-align: center;
    img {
      vertical-align: middle;
      width: 70%;
    }
  }
}

</style>