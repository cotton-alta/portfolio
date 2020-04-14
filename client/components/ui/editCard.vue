<template>
  <div class="card-wrapper">
    <div class="card-title">
      {{ card.Title }}
    </div>
    <!-- /.card-title -->
    <div class="card-content">
      {{ card.Timestamp }}
    </div>
    <!-- /.card-content -->
    <deleteButton class="card-button" :title="card.Title" />
  </div>
  <!-- /.card-wrapper -->
</template>

<script>
import deleteButton from "~/components/ui/deleteButton.vue"

export default {
  props: ["card"],
  components: { deleteButton }
}
</script>

<style lang="scss" scoped>

.card {
  &-wrapper {
    width: 80%;
    // height: 100px;
    margin: 0 auto;
    background: $black-background;
    border-radius: 5px;
    border: 2px solid $menu-background;
  }
  &-title {
    color: $menu-background;
    font-size: calc(20px + 0.4vw);
    font-weight: bolder;
    margin: 10px 0px 10px 10px;
  }
  &-content {
    line-height: 30px;
    color: $menu-background;
    font-size: calc(15px + 0.3vw);
    font-weight: bolder;
    margin: 10px 0px 10px 10px;
  }
  &-button {
    margin: 10px 0px 10px 10px;
  }
}

</style>