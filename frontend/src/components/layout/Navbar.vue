<template>
  <nav
    class="navbar navbar-vertical fade-in fixed-left navbar-expand-md navbar-dark navbar-vibrant"
    style="overflow-y: overlay"
    id="sidebar"
    @click.self="expandNavbar()"
    v-if="!hideNavbar"
  >

    <div
      style="position: inherit; top: 50%; left: 10px; display: none"
      @click="expandNavbar()"
      id="collapse-nav-chevron"
    >
      <i
        class="fa fa-chevron-right navbar-chevron"
        style="font-size: 40px; color: white; position: relative;"
      ></i>
    </div>

    <div class="container-fluid navbar-contents">

      <!-- Toggler -->
      <button
        class="navbar-toggler" type="button" data-toggle="collapse" data-target="#sidebarCollapse"
        aria-controls="sidebarCollapse" aria-expanded="false" aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <!-- Brand -->
      <!--      <router-link class="ml-3 mt-3 d-none d-lg-block" to="/">-->
      <!--        <img-->
      <!--          src="~@/assets/img/eqemu-logo-1.png"-->
      <!--          class="navbar-brand-img mx-auto d-none d-sm-block mb-3" alt="..."-->
      <!--          style="max-height: 6rem"-->
      <!--        >-->
      <!--      </router-link>-->


      <router-link class="ml-3 mt-3" to="/">
        <h1 class="text-center eq-header small-mobile">
          Spire
        </h1>

      </router-link>

      <!--      <hr class="dropdown-divider">-->


      <!--      <hr class="dropdown-divider">-->

      <!--      <h4 class=" text-center menuetto-header mt-2" style="font-size: 34px">-->
      <!--        Tools-->
      <!--      </h4>-->

      <!-- User (xs) -->
      <div class="navbar-user d-md-none">

        <!-- Dropdown -->
        <div class="dropdown">

          <!-- Toggle -->
          <a
            href="#" id="sidebarIcon" class="dropdown-toggle" role="button" data-toggle="dropdown" aria-haspopup="true"
            aria-expanded="false"
          >
            <div class="avatar avatar-sm avatar-online">
              <img
                :src="user ? user.avatar : require('@/assets/img/eqemu-avatar.png')"
                class="avatar-img rounded-circle" alt="..."
              >
            </div>
          </a>

          <navbar-dropdown-menu menu-right="1"/>

        </div>

      </div>

      <!-- Collapse -->
      <div class="collapse navbar-collapse" id="sidebarCollapse">

        <!-- Form -->
        <!--        <form class="mt-4 mb-3 d-md-none">-->
        <!--          <div class="input-group input-group-rounded input-group-merge">-->
        <!--            <input type="search" class="form-control form-control-rounded form-control-prepended" placeholder="Search"-->
        <!--                   aria-label="Search">-->
        <!--            <div class="input-group-prepend">-->
        <!--              <div class="input-group-text">-->
        <!--                <span class="fe fe-search"></span>-->
        <!--              </div>-->
        <!--            </div>-->
        <!--          </div>-->
        <!--        </form>-->

        <!-- Heading -->
        <h6 class="navbar-heading mt-3">
          Tools
        </h6>

        <!-- Navigation -->
        <ul class="navbar-nav mb-md-4">
          <nav-section-component :config="botNav" v-if="alphaEnabled"/>

          <nav-section-component :config="calculatorNav"/>

          <li class="nav-item">
            <router-link class="nav-link" to="/client-files">
              <i class="ra ra-cycle mr-1"></i> Client Files
              <b-badge class="ml-3" variant="primary">NEW!</b-badge>
            </router-link>
          </li>

          <li class="nav-item">
            <router-link class="nav-link " to="/items">
              <i class="ra ra-relic-blade mr-1"></i> Items
              <b-badge class="ml-3" variant="primary">NEW!</b-badge>
            </router-link>
          </li>

          <nav-section-component :config="npcNav" v-if="alphaEnabled"/>

          <li class="nav-item">
            <router-link class="nav-link " to="/quest-api-explorer">
              <i class="ra ra-compass mr-1"></i> Quest API Explorer
            </router-link>
          </li>

          <li class="nav-item">
            <router-link class="nav-link " to="/strings-database">
              <i class="ra  ra-scroll-unfurled mr-1"></i> Strings DB
              <b-badge class="ml-3" variant="primary">NEW!</b-badge>
            </router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link " to="/spells">
              <i class="ra ra-book mr-1"></i> Spells
              <b-badge class="ml-3" variant="primary">NEW!</b-badge>
            </router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link " to="/tasks">
              <i class="ra ra-zebra-shield mr-1"></i> Tasks
              <b-badge class="ml-3" variant="primary">BETA</b-badge>
              <b-badge class="ml-3" variant="primary">NEW!</b-badge>
            </router-link>
          </li>
          <nav-section-component :config="viewerNav"/>

          <li class="nav-item" v-if="alphaEnabled">
            <router-link class="nav-link " to="/zones">
              <i class="ra ra-wooden-sign mr-2"></i> Zones
              <b-badge class="ml-3" variant="primary">ALPHA</b-badge>
              <b-badge class="ml-3" variant="primary">NEW!</b-badge>
            </router-link>
          </li>

        </ul>

        <!-- Heading -->
        <h6 class="navbar-heading">
          Spire Docs
        </h6>

        <!-- Navigation -->
        <ul class="navbar-nav mb-md-4">

          <nav-section-component :config="spireApiNav"/>

          <!-- Components -->
          <li class="nav-item">
            <a
              :class="'nav-link collapse ' + (hasRoute('components') ? 'active' : 'collapsed')"
              href="#sidebarComponents" data-toggle="collapse" role="button"
              aria-expanded="false" aria-controls="sidebarComponents"
            >
              <i class="ra ra-burst-blob mr-1"></i> Components
            </a>
            <div :class="'collapse ' + (hasRoute('components') ? 'show' : '')" id="sidebarComponents">
              <ul class="nav nav-sm flex-column">
                <li v-for="nav in componentNavs">
                  <router-link class="nav-link" :to="nav.to">{{ nav.title }}</router-link>
                </li>
              </ul>
            </div>
          </li>

        </ul>

        <h6 class="navbar-heading" v-if="appVersion">
          Version ({{ appEnv }}) {{ appVersion }}
        </h6>

        <!-- Push content down -->
        <div class="mt-auto"></div>

        <!-- User (md) -->
        <div class="navbar-user d-none d-md-flex" id="sidebarUser">
          <navbar-user-settings-cog/>

          <!-- Dropup -->
          <div class="dropup">

            <!-- Toggle -->
            <a
              href="#" id="sidebarIconCopy" class="dropdown-toggle" role="button" data-toggle="dropdown"
              aria-haspopup="true" aria-expanded="false"
            >
              <div class="avatar avatar-sm avatar-online">
                <img
                  :src="user ? user.avatar : require('@/assets/img/eqemu-avatar.png')"
                  class="avatar-img rounded-circle" alt="..."
                >
              </div>
            </a>

            <!-- Menu -->
            <navbar-dropdown-menu/>

          </div>

          <!-- Icon -->
          <a
            href="javascript:void(0)"
            class="navbar-user-link"
            @click="collapseNavbar"
          >
              <span class="icon">
                <i class="fe fe-menu"></i>
              </span>
          </a>

        </div>

        <!-- Active Database Connection Status -->
        <db-connection-status-pill/>

      </div> <!-- / .navbar-collapse -->

    </div>
  </nav>
</template>

<script>

import {App}                  from "@/constants/app";
import NavbarDropdownMenu     from "@/components/layout/NavbarDropdownMenu";
import NavbarUserSettingsCog  from "@/components/layout/NavbarUserSettingsCog";
import UserContext            from "@/app/user/UserContext";
import NavSectionComponent    from "@/components/layout/NavSectionComponent";
import {ROUTE}                from "@/routes";
import {EventBus}             from "@/app/event-bus/event-bus";
import {AppEnv}               from "@/app/env/app-env";
import {Navbar}               from "@/app/navbar";
import DbConnectionStatusPill from "@/components/DbConnectionStatusPill";
import {SpireApi}             from "@/app/api/spire-api";

export default {
  components: { DbConnectionStatusPill, NavSectionComponent, NavbarDropdownMenu, NavbarUserSettingsCog },
  data() {
    return {
      backendBaseUrl: "",
      user: null,
      hideNavbar: false,
      alphaEnabled: App.ALPHA_TOOLS_ENABLED,
      appEnv: AppEnv.getEnv(),
      appVersion: AppEnv.getVersion(),
      appFeatures: AppEnv.getFeatures(),
      botNav: {
        label: "Bots",
        labelIcon: "ra ra-reactor mr-1",
        routePrefixMatches: ["bot"],
        navs: [
          {
            title: "Spells",
            to: ROUTE.BOT_SPELLS_EDIT,
            icon: "ra ra-regeneration mr-1",
            isAlpha: true,
            isNew: true,
            routes: ['bot-spells']
          },
        ]
      },
      npcNav: {
        label: "NPCs",
        labelIcon: "ra ra-dragon mr-1",
        routePrefixMatches: ["npc", "merchant", "loot"],
        navs: [
          {
            title: "Emotes",
            to: ROUTE.NPC_EMOTES_EDIT,
            icon: "ra ra-death-skull mr-1",
            isAlpha: true,
            isNew: true,
            routes: ['npc-emotes']
          },
          {
            title: "Merchants",
            to: ROUTE.MERCHANTS,
            icon: "ra ra-emerald mr-1",
            isAlpha: true,
            isNew: true,
            routes: ['merchant', 'merchants']
          },
          {
            title: "Spells",
            to: ROUTE.NPC_SPELLS_EDIT,
            icon: "ra ra-flame-symbol mr-1",
            isAlpha: true,
            isNew: true,
            routes: ['npc-spells']
          },
          // {
          //   title: "Loot",
          //   to: ROUTE.LOOT,
          //   icon: "ra ra-sword mr-1",
          //   isAlpha: true,
          //   isNew: true,
          //   routes: ['loot']
          // },
        ]
      },
      viewerNav: {
        label: "Viewers",
        labelIcon: "ra ra-bleeding-eye mr-1",
        routePrefixMatch: "viewer",
        navs: [
          { title: "Race Viewer", to: ROUTE.RACE_VIEWER, icon: "ra ra-monster-skull mr-1" },
          { title: "Item Model Viewer", to: ROUTE.ITEM_VIEWER, icon: "ra ra-crossed-swords mr-1" },
          { title: "Item Icon Viewer", to: ROUTE.ITEM_ICON_VIEWER, icon: "ra ra-burning-eye mr-1" },
          {
            title: "Player Animations",
            to: ROUTE.PLAYER_ANIMATION_VIEWER,
            icon: "ra ra-player-dodge mr-1",
            isNew: true
          },
          { title: "Emitter Viewer", to: ROUTE.EMITTER_VIEWER, icon: "ra  ra-droplet-splash mr-1", isNew: true },
          { title: "Spell Animations", to: ROUTE.SPELL_ANIMATION_VIEWER, icon: "ra ra-dragon mr-1", isNew: true }
        ]
      },
      spireApiNav: {
        label: "Spire API Docs",
        labelIcon: "ra ra-book mr-1",
        routePrefixMatches: ["swagger", "model-viewer"],
        navs: [
          { title: "Swagger API", to: SpireApi.getBasePath() + '/swagger/index.html', icon: "ra ra-monster-skull mr-1" },
          { title: "Model Relationship Explorer", to: ROUTE.API_MODEL_RELATIONSHIP_EXPLORER, icon: "ra ra-kaleidoscope mr-1" },
        ]
      },
      componentNavs: [
        { title: "Facial Appearance", to: "/components#facial-appearance" },
        { title: "Progress Bars", to: "/components#progress-bars" },
        { title: "Page Headers", to: "/components#page-headers" },
        { title: "Range Visual", to: "/components#range-visual" },
        { title: "Item Preview", to: "/components#item-preview" },
        { title: "Spell Icons", to: "/components#spell-icons" },
        { title: "Spell Preview", to: "/components#spell-preview" },
        { title: "NPC Card Preview", to: "/components#npc-card-preview" },
        { title: "NPC Special Abilities", to: "/components#npc-special-abilities" },
        { title: "Table", to: "/components#table" },
        { title: "Tabs", to: "/components#tabs" },
        { title: "Form Elements", to: "/components#form-elements" },
        { title: "Windows", to: "/components#windows" }
      ],
      calculatorNav: {
        label: "Calculators",
        labelIcon: "ra ra-cog mr-1",
        routePrefixMatch: "calculator",
        navs: [
          { title: "Race Calculator", to: "/calculators#race-bitmask-calculator", icon: "ra ra-eye-monster mr-1" },
          { title: "Class Calculator", to: "/calculators#class-bitmask-calculator", icon: "ra ra-lion mr-1" },
          { title: "Deity Calculator", to: "/calculators#deity-bitmask-calculator", icon: "ra ra-venomous-snake mr-1" },
          {
            title: "Client Version Calculator",
            to: "/calculators#client-version-calculator",
            icon: "ra ra-lever mr-1"
          },
          {
            title: "Expansions Calculator",
            to: "/calculators#expansions-bitmask-calculator",
            icon: "ra ra-lever mr-1"
          },
          { title: "Augment Type Calculator", to: "/calculators#augment-type-calculator", icon: "ra ra-sapphire mr-1" },
          {
            title: "Inventory Slot Calculator",
            to: "/calculators#inventory-slot-calculator",
            icon: "ra ra-eye-shield mr-1"
          },
          { title: "Special Abilities Calculator", to: "/calculators#npc-special-abilities", icon: "ra ra-lion mr-1" },
        ]
      },


    }
  },
  created() {
    EventBus.$on("HIDE_NAVBAR", this.toggleNavbarCollapse);
    EventBus.$on("APP_ENV_LOADED", this.handleAppEnvLoaded);
    EventBus.$on("ALPHA_ENABLED", this.handleAppAlphaEnabled);
  },
  destroyed() {
    EventBus.$off("HIDE_NAVBAR", this.toggleNavbarCollapse);
    EventBus.$off("APP_ENV_LOADED", this.handleAppEnvLoaded);
    EventBus.$off("ALPHA_ENABLED", this.handleAppAlphaEnabled);
  },

  async mounted() {
    this.backendBaseUrl = App.BACKEND_BASE_URL
    this.user           = await UserContext.getUser()
  },

  methods: {
    handleAppAlphaEnabled() {
      this.alphaEnabled = App.ALPHA_TOOLS_ENABLED
    },
    handleAppEnvLoaded() {
      this.appEnv      = AppEnv.getEnv();
      this.appVersion  = AppEnv.getVersion();
      this.appFeatures = AppEnv.getFeatures();
    },
    expandNavbar() {
      Navbar.expand()
    },
    collapseNavbar() {
      Navbar.collapse()
    },
    toggleNavbarCollapse() {
      Navbar.toggleCollapse()
    },
    hideNavBar() {
      this.hideNavbar = !this.hideNavbar
    },
    hasRoute: function (partial) {
      return (this.$route.path.indexOf(partial) > -1)
    },
    hideNavbarAfterClick() {
      const sidebar = document.getElementById("sidebarCollapse")
      if (sidebar) {
        sidebar.classList.remove("show");
      }
    },
  },
  watch: {
    $route(to, from) {
      this.hideNavbarAfterClick()
    }
  }
}
</script>

<style scoped>
.navbar-nav .nav-link > .fe {
  min-width: 1.25rem;
}

.navbar-dark.navbar-vibrant {
  background-image: linear-gradient(to bottom right, rgb(21 21 21 / 90%), rgb(191 126 70 / 90%)), url(~@/assets/img/eq-wallpaper-1.jpg);
}

@media only screen and (max-device-width: 640px) {
  .small-mobile {
    font-size: 40px !important;
  }
}

.connection-status-box {
  display: block; /* or inline-block */
  text-overflow: ellipsis;
  word-wrap: break-word;
  overflow: hidden;
}

</style>
