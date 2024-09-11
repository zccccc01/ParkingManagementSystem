import sanitizeHtml from 'sanitize-html';

const sanitizeInput = (input) => {
  return sanitizeHtml(input, {
    allowedTags: [],
    allowedAttributes: {},
  });
};

export default sanitizeInput;
