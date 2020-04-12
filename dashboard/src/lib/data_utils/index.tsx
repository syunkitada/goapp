function getIndexDataFromState(state, index) {
    let service: any = null;
    let serviceState = state.service;
    if (serviceState.projectName) {
        service =
            serviceState.projectServiceMap[serviceState.projectName][
                serviceState.serviceName
            ];
    } else {
        service = serviceState.serviceMap[serviceState.serviceName];
    }
    return service.Data[index.DataKey];
}

export default {
    getIndexDataFromState
};
