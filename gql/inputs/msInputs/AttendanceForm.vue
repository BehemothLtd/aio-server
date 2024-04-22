<template>
  <div class="row">
    <div class="card">
      <div class="card-body">
        <div class="row mb-3">
          <FormValidator
            label="User"
            name="UserId"
            required
          >
            <VSelect
              v-model="attendance.userId"
              class="w-100"
              :options="userOptions"
              :reduce="(option) => option.value"
            >
            </VSelect>
          </FormValidator>

          <FormValidator
            class="mt-2"
            name="CheckinAt"
            label="Check In"
            required
          >
            <DatePicker
              v-model="attendance.checkinAt"
              :format="formatDate"
              class="date-picker"
              placeholder="Select Date"
            ></DatePicker>
          </FormValidator>

          <FormValidator
            class="mt-2"
            name="CheckoutAt"
            label="Check Out"
          >
            <DatePicker
              v-model="attendance.checkoutAt"
              class="date-picker"
              :format="formatDate"
              placeholder="Select Date"
            ></DatePicker>
          </FormValidator>
        </div>

        <div class="row text-end">
          <div>
            <button
              type="button"
              class="btn btn-primary w-md"
              @click="submitForm"
            >
              Submit
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  userOptions: Array,
  default: () => [],
});

const attendance = defineModel();
const emits = defineEmits(["create", "update"]);

const formatDate = "dd-MM-yyyy HH:mm";

function submitForm() {
  if (attendance.value.id) emits("update", attendance.value.id);
  else emits("create");
}
</script>
