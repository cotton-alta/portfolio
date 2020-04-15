<template>
  <div class="footer-wrapper">
    <div class="footer-image">
      <div><a href="https://github.com/cotton-alta"><img src="/github.png" /></a></div>
      <div><a href="https://twitter.com/cotton_alta_"><img src="/twitter.png" /></a></div>
      <div><a href="mailto:chinoknct@gmail.com"><img src="/mail.png" /></a></div>
    </div>
    <div class="copyright">
      <p>&copy;2020 cotton</p>
    </div>
    <!-- /.footer-image -->
  </div>
</template>

<script>
export default {
  
}
</script>

<style lang="scss" scoped>

.footer {
  &-wrapper {
    line-height: 30px;
    width: 100%;
    height: 150px;
    background: $black-background;
    display: block;
    position: relative;
  }
  &-image {
    width: 30%;
    height: 50px;
    padding-top: 20px;
    margin: 0 auto;
    justify-content: center;
    display: flex;
    flex-direction: row;
    img {
      width: 2.5rem;
      margin: 0px 25px;
    }
  }
}

.copyright {
  width: 100px;
  margin: 0 auto;
  position: absolute;
  left: calc(50% - 50px);
  bottom: 10px;
  color: $string-color;
}

</style>