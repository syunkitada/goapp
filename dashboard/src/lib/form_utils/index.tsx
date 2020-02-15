import logger from '../logger';

function getSearchQueries() {
  const queryStr = decodeURIComponent(location.search);
  let searchQueries = {};
  try {
    const value = queryStr.match(new RegExp('[?&]q=({.*?})(&|$|#)'));
    if (value) {
      searchQueries = JSON.parse(value[1]);
    }
  } catch (e) {
    logger.warning('Ignored failed parse', queryStr);
  }
  return searchQueries;
}

function mergeDefaultInputsToFormData(index, rawData, formData) {
  const newFormData = {};
  const inputErrorMap = {};
  if (index.Inputs) {
    // Validate
    // フォーム入力がなく、デフォルト値がある場合はセットする
    for (let i = 0, len = index.Inputs.length; i < len; i++) {
      const input = index.Inputs[i];
      let value = formData[input.Name];
      switch (input.Type) {
        case 'Text':
          if (input.Require) {
            if (!value || value === '') {
              inputErrorMap[input.Name] = {
                error: 'This is required',
                type: input.Kind,
              };
            }
            newFormData[input.Name] = {value};
          }
          break;

        case 'Select':
          if (!value) {
            if (input.Default) {
              value = input.Default;
            } else {
              let options = input.Options;
              if (!options) {
                options = [];
                const d = rawData[input.DataKey];
                if (d) {
                  for (let j = 0, l = d.length; j < l; j++) {
                    options.push(d[j].Name);
                  }
                } else {
                  options.push('');
                }
              }
              value = options[0];
            }
          }
          newFormData[input.Name] = value;
          break;
        case 'DateTime':
          newFormData[input.Name] = value;
          break;
        default:
          break;
      }
    }
  }
  return {newFormData, inputErrorMap};
}

export default {
  getSearchQueries,
  mergeDefaultInputsToFormData,
};
