const {Firestore} = require('@google-cloud/firestore');

exports.getAutomoveis = async function (req, res) {
  res.setHeader("Access-Control-Allow_Origin", "*");

  const placa = req.query.placa
  console.log('Placa passada por url: ', placa);

  const firestore = new Firestore();

  let query = firestore.collection('automoveis').where('Placa', '==', placa);

  let automoveis = []

  await query.get()
  .then(querySnapshot => {
    querySnapshot.docs.forEach(documentSnapshot => {
      const {Marca,Modelo,Cor,Preco,Placa} = documentSnapshot._fieldsProto
      let automovel = {
        Marca: Marca.stringValue,
        Modelo: Modelo.stringValue,
        Cor: Cor.stringValue,
        Preco: Preco.stringValue,
        Placa: Placa.stringValue
      };
      automoveis.push(automovel)
    })
  }).catch(err => {
    console.log('Erro ao fazer a query: ', err);
    res.status(500).end(err);
  });

  res.status(200).end(JSON.stringify(automoveis));
}

