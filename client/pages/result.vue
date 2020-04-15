<template>
  <div class="container">
    <div class="text"><span>お問い合わせありがとうございます。</span></div>
    <div class="text"><span>数日中に折り返しご連絡させていただきますので今しばらくお待ちくださいませ。</span></div>
  </div>
  <!-- /.container -->
</template>

<script>
export default {
  
}
</script>

<style lang="scss" scoped>
.container {
  display: block;
  width: 90%;
  max-width: 500px;
  margin: 0 auto;
  padding-top: 20%;
  min-height: calc(100vh - 50px);
  @include mq {
    min-height: calc(100vh - 80px);
  }
  span {
    width: 100%;
    color: $string-color;
    font-size: calc(20px + 0.5vw);
    line-height: 30px;
  }
}

.text {
  margin-bottom: 20px;
}
</style>