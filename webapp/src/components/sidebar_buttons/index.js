// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import {list, updateRhsState, telemetry} from '../../actions';

import SidebarButtons from './sidebar_buttons.jsx';

function mapStateToProps(state) {
    return {
        issues: state['plugins-com.schat.plugin-todo-next'].issues,
        inIssues: state['plugins-com.schat.plugin-todo-next'].inIssues,
        outIssues: state['plugins-com.schat.plugin-todo-next'].outIssues,
        showRHSPlugin: state['plugins-com.schat.plugin-todo-next'].rhsPluginAction,
    };
}

function mapDispatchToProps(dispatch) {
    return {
        actions: bindActionCreators({
            list,
            updateRhsState,
            telemetry,
        }, dispatch),
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(SidebarButtons);