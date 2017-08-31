package product 

  import (
"github.com/hexya-erp/hexya/hexya/models/types"

) 

 func init() { 

 

pool.ProductUomCateg().DeclareModel()
pool.ProductUomCateg().AddCharField("Name", models.StringFieldParams{String :"Name", Required : true, Translate: true})


pool.ProductUom().DeclareModel()
pool.ProductUom().AddCharField("Name", models.StringFieldParams{String :"Unit of Measure", Required : true, Translate: true})
pool.ProductUom().AddMany2OneField("Category",models.ForeignKeyFieldParams{String :"Category" , RelationModel: pool.ProductUomCateg() , JSON : "category_id", Required : true, OnDelete : models.Cascade , Help :"Conversion between Units of Measure can only occur if they belong to the same category. The conversion will be made based on the ratios."})
pool.ProductUom().AddFloatField("Factor", models.FloatFieldParams{String :"Ratio", Default: func(models.Environment, models.FieldMap) interface{} {return 1.0}, Required : true})
pool.ProductUom().AddFloatField("FactorInv", models.FloatFieldParams{String :"Bigger Ratio", Required : true ,Help :"'How many times this Unit of Measure is bigger than the reference Unit of Measure in this category: 1 * "})
pool.ProductUom().AddFloatField("Rounding", models.FloatFieldParams{String :"Rounding Precision", Default: func(models.Environment, models.FieldMap) interface{} {return 0.01}, Required : true ,Help :"The computed quantity will be a multiple of this value.  Use 1.0 for a Unit of Measure that cannot be further split, such as a piece."})
pool.ProductUom().AddBooleanField("Active", models.SimpleFieldParams{String :"Active", Default: func(models.Environment, models.FieldMap) interface{} {return true} ,Help :"Uncheck the active field to disable a unit of measure without deleting it."})
pool.ProductUom().AddSelectionField("UomType", models.SelectionFieldParams{String :"Type", Selection : types.Selection{
"bigger" : "Bigger than the reference Unit of Measure",
"reference" : "Reference Unit of Measure for this category",
"smaller" : "Smaller than the reference Unit of Measure",
}, Default: func(models.Environment, models.FieldMap) interface{} {return "reference"}, Required: true})
pool.ProductUom().AddSQLConstraint("FactorGtZero" , "CHECK (factor!=0)" , "The conversion ratio for a unit of measure cannot be 0!")
pool.ProductUom().AddSQLConstraint("RoundingGtZero" , "CHECK (rounding>0)" , "The rounding precision must be greater than 0!")
pool.ProductUom().Methods().ComputeFactorInv().DeclareMethod(
`ComputeFactorInv` ,
func (rs pool.ProductUomSet){
  //@api.depends('factor')
  /*def _compute_factor_inv(self):
        self.factor_inv = self.factor and (1.0 / self.factor) or 0.0

    */})
pool.ProductUom().Methods().OnchangeUomType().DeclareMethod(
`OnchangeUomType` ,
func (rs pool.ProductUomSet){
  //@api.onchange('uom_type')
  /*def _onchange_uom_type(self):
        if self.uom_type == 'reference':
            self.factor = 1

    */})
pool.ProductUom().Methods().Create().DeclareMethod(
`Create` ,
func (rs pool.ProductUomSet , args struct{Values interface{}
}){
  //@api.model
  /*def create(self, values):
        if 'factor_inv' in values:
            factor_inv = values.pop('factor_inv')
            values['factor'] = factor_inv and (1.0 / factor_inv) or 0.0
        return super(ProductUoM, self).create(values)

    */})
pool.ProductUom().Methods().Write().DeclareMethod(
`Write` ,
func (rs pool.ProductUomSet , args struct{Values interface{}
}){
  //@api.multi
  /*def write(self, values):
        if 'factor_inv' in values:
            factor_inv = values.pop('factor_inv')
            values['factor'] = factor_inv and (1.0 / factor_inv) or 0.0
        return super(ProductUoM, self).write(values)

    */})
pool.ProductUom().Methods().NameCreate().DeclareMethod(
`NameCreate` ,
func (rs pool.ProductUomSet , args struct{}){
  //@api.model
  /*def name_create(self, name):
        """ The UoM category and factor are required, so we'll have to add temporary values
        for imported UoMs """
        values = {
            self._rec_name: name,
            'factor': 1
        }
        # look for the category based on the english name, i.e. no context on purpose!
        # TODO: should find a way to have it translated but not created until actually used
        if not self._context.get('default_category_id'):
            EnglishUoMCateg = self.env['product.uom.categ'].with_context({})
            misc_category = EnglishUoMCateg.search([('name', '=', 'Unsorted/Imported Units')])
            if misc_category:
                values['category_id'] = misc_category.id
            else:
                values['category_id'] = EnglishUoMCateg.name_create('Unsorted/Imported Units')[0]
        new_uom = self.create(values)
        return new_uom.name_get()[0]

    */})
pool.ProductUom().Methods().ComputeQuantity().DeclareMethod(
`ComputeQuantity` ,
func (rs pool.ProductUomSet , args struct{ToUnit interface{}
Round interface{}
RoundingMethod interface{}
}){
  //@api.multi
  /*def _compute_quantity(self, qty, to_unit, round=True, rounding_method='UP'):
        if not self:
            return qty
        self.ensure_one()
        if self.category_id.id != to_unit.category_id.id:
            if self._context.get('raise-exception', True):
                raise UserError(_('Conversion from Product UoM %s to Default UoM %s is not possible as they both belong to different Category!.') % (self.name, to_unit.name))
            else:
                return qty
        amount = qty / self.factor
        if to_unit:
            amount = amount * to_unit.factor
            if round:
                amount = tools.float_round(amount, precision_rounding=to_unit.rounding, rounding_method=rounding_method)
        return amount

    */})
pool.ProductUom().Methods().ComputePrice().DeclareMethod(
`ComputePrice` ,
func (rs pool.ProductUomSet , args struct{Price interface{}
ToUnit interface{}
}){
  //@api.multi
  /*def _compute_price(self, price, to_unit):
        self.ensure_one()
        if not self or not price or not to_unit or self == to_unit:
            return price
        if self.category_id.id != to_unit.category_id.id:
            return price
        amount = price * self.factor
        if to_unit:
            amount = amount / to_unit.factor
        return amount
*/})
 
 }