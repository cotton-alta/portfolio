<template>
  <div class="container">
    <topTitle />
    <topButton />
  </div>
  <!-- /.container -->
</template>

<script>
import topTitle from "~/components/layouts/topTitle.vue"
import topButton from "~/components/ui/button.vue"

export default {
  components: {
    topTitle,
    topButton
  }
}
</script>

<style>
.container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
</style>
