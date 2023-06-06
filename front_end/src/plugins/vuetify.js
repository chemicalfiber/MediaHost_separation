import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';
import colors from 'vuetify/lib/util/colors'

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        themes: {
            light: {
                primary: colors.orange, // #FF9800
                secondary: colors.orange.lighten2, // #FFB74D
                accent: colors.orange.accent4, // ##FF6D00
            },
        },
    },
});
