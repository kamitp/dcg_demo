db = db.getSiblingDB('recipes_service');
db.createCollection("idgenetor");
db.createCollection("recipes");
db.idgenetor.insert({_id:"recipeid", sequence_value:0});
db.system.js.save(
  {
    _id : "getNextId" ,
    value : function (recipeid){ 
        var document = db.idgenetor.findAndModify({
        query:{_id: recipeid },
        update: {$inc:{sequence_value:1}},
        new:true});
        return document.sequence_value; 
       }
  }
);
db.loadServerScripts();
db.recipes.insert({
  "_id":getNextId("recipeid"),
  "name":"Cucumber Pasta"
});
