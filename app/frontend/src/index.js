require("./less/theme.less")

import UIkit from 'uikit';
import Icons from 'uikit/dist/js/uikit-icons';
import dragula from 'dragula/dragula';

// loads the Icon plugin
UIkit.use(Icons);

// components can be called from the imported UIkit reference
UIkit.notification('Hello world.');


function $ (id) {
    return document.getElementById(id);
  }
dragula([$('left-defaults'), $('right-defaults')]);