

var xml = "<?xml version=\"1.0\" encoding=\"UTF-8\" ?><document><alertTemplate> \
 <title>TVJS sample</title> \
 <description>this screen presented by TVJS javascript</description> \
 </alertTemplate></document>";

var parser = new DOMParser();
var alertDoc = parser.parseFromString(xml, "application/xml");
navigationDocument.presentModal(alertDoc);
