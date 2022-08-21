<template>
  <div class="min-h-full">
    <div class="bg-white shadow-sm relative z-10">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="hidden sm:-my-px sm:ml-6 sm:flex sm:space-x-8" :key="updateNavBar">
              <a
                v-for="item in navigation"
                :key="item.name"
                :href="item.href"
                :class="[
                  item.current
                    ? 'border-blue-500 text-gray-900'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
                  'inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium',
                ]"
                :aria-current="item.current ? 'page' : undefined"
                >{{ item.name }}</a
              >
              <div v-if="grid && curNavElement != 0" class="inline-flex items-center px-1 pt-1 text-sm font-medium">
                {{ grid.data.name }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="sm:hidden">
        <div class="pt-2 pb-3 space-y-1">
          <a
            v-for="item in navigation"
            :key="item.name"
            :href="item.href"
            :class="[
              item.current
                ? 'bg-blue-50 border-blue-500 text-blue-700'
                : 'border-transparent text-gray-600 hover:bg-gray-50 hover:border-gray-300 hover:text-gray-800',
              'block pl-3 pr-4 py-2 border-l-4 text-base font-medium',
            ]"
            :aria-current="item.current ? 'page' : undefined"
          >
            {{ item.name }}
          </a>
        </div>
      </div>
    </div>

    <div>
      <router-view />
    </div>
  </div>
</template>

<script lang="ts">
import { emitter } from "@/components/util/Emittery";
import GridData from "@/models/Data/GridData";
import GridModel from "@/models/GridModel";
import { Disclosure, DisclosureButton, DisclosurePanel } from "@headlessui/vue";
import { defineComponent, onMounted, ref } from "vue";

export default defineComponent({
  setup(props, ctx) {
    const grid = ref(GridData.grid);
    const updateNavBar = ref(0);
    const curNavElement = ref(0);

    const navigation = ref([
      { name: "Overview", href: "/", current: true },
      { name: "Workflow Editor", href: "", current: false },
      { name: "Dashboard", href: "", current: false },
    ]);

    emitter.on("UpdateNavBar", ([index, id]) => {
      console.log("Update", GridData.grid);
      grid.value = GridData.grid;

      // Set Active Navigation
      curNavElement.value = index;
      let curNav = navigation.value.find((nav) => nav.current);
      if (curNav) {
        curNav.current = false;
      }
      navigation.value[index].current = true;

      if (id) {
        navigation.value[1].href = "/workflow-editor/" + id;
        navigation.value[2].href = "/dashboard/" + id;
      }

      updateNavBar.value++;
    });

    return {
      navigation,
      grid,
      updateNavBar,
      curNavElement,
    };
  },
});
</script>