package product 

  import (

) 

 func init() { 

 

pool.ProductAttribute().DeclareModel()
pool.ProductAttribute().AddCharField("Name", models.StringFieldParams{String :"Name", Required: true, Translate: true})
pool.ProductAttribute().AddOne2ManyField("ValueIds", models.ReverseFieldParams{String :"Values" ,RelationModel : pool.ProductAttributeValue() ,ReverseFK : "AttributeId", NoCopy: false})
pool.ProductAttribute().AddIntegerField("Sequence", models.SimpleFieldParams{String :"Sequence" ,Help :"Determine the display order"})
pool.ProductAttribute().AddOne2ManyField("AttributeLineIds", models.ReverseFieldParams{String :"Lines')" ,RelationModel : pool.ProductAttributeLine() ,ReverseFK : "AttributeId"})
pool.ProductAttribute().AddBooleanField("CreateVariant", models.SimpleFieldParams{String :"CreateVariant", Default: func(models.Environment, models.FieldMap) interface{} {return true} ,Help :"Check this if you want to create multiple variants for this attribute."})


pool.ProductAttributevalue().DeclareModel()
pool.ProductAttributevalue().AddCharField("Name", models.StringFieldParams{String :"Value", Required: true, Translate: true})
pool.ProductAttributevalue().AddIntegerField("Sequence", models.SimpleFieldParams{String :"Sequence" ,Help :"Determine the display order"})
pool.ProductAttributevalue().AddMany2OneField("AttributeId",models.ForeignKeyFieldParams{String :"Attribute" , RelationModel: pool.ProductAttribute(), OnDelete : models.Cascade, Required: true})
pool.ProductAttributevalue().AddMany2ManyField("ProductIds", models.Many2ManyFieldParams{String :"Variants" , RelationModel: pool.ProductProduct()})
pool.ProductAttributevalue().Fields().ProductIds().RevokeAccess(security.GroupEveryone, security.Write)
pool.ProductAttributevalue().AddFloatField("PriceExtra", models.FloatFieldParams{String :"Attribute Price Extra", Compute: "ComputePriceExtra", Default: func(models.Environment, models.FieldMap) interface{} {return 0.0}})
pool.ProductAttributevalue().AddOne2ManyField("PriceIds", models.ReverseFieldParams{String :"Attribute Prices" ,RelationModel : pool.ProductAttributePrice() ,ReverseFK : "ValueId"})
pool.ProductAttributevalue().AddSQLConstraint("ValueCompanyUniq" , "Unique (name,AttributeId)" , "This attribute value already exists !")
pool.ProductAttributevalue().Methods().ComputePriceExtra().DeclareMethod(
`ComputePriceExtra` ,
func (rs pool.ProductAttributevalueSet){
  //@api.one
  /*def _compute_price_extra(self):
        if self._context.get('active_id'):
            price = self.price_ids.filtered(lambda price: price.product_tmpl_id.id == self._context['active_id'])
            self.price_extra = price.price_extra
        else:
            self.price_extra = 0.0

    */})
pool.ProductAttributevalue().Methods().SetPriceExtra().DeclareMethod(
`SetPriceExtra` ,
func (rs pool.ProductAttributevalueSet){
  /*def _set_price_extra(self):
        if not self._context.get('active_id'):
            return

        AttributePrice = self.env['product.attribute.price']
        prices = AttributePrice.search([('value_id', 'in', self.ids), ('product_tmpl_id', '=', self._context['active_id'])])
        updated = prices.mapped('value_id')
        if prices:
            prices.write({'price_extra': self.price_extra})
        else:
            for value in self - updated:
                AttributePrice.create({
                    'product_tmpl_id': self._context['active_id'],
                    'value_id': value.id,
                    'price_extra': self.price_extra,
                })

    */})
pool.ProductAttributevalue().Methods().NameGet().DeclareMethod(
`NameGet` ,
func (rs pool.ProductAttributevalueSet){
  //@api.multi
  /*def name_get(self):
        if not self._context.get('show_attribute', True):  # TDE FIXME: not used
            return super(ProductAttributevalue, self).name_get()
        return [(value.id, "%s: %s" % (value.attribute_id.name, value.name)) for value in self]

    */})
pool.ProductAttributevalue().Methods().Unlink().DeclareMethod(
`Unlink` ,
func (rs pool.ProductAttributevalueSet){
  //@api.multi
  /*def unlink(self):
        linked_products = self.env['product.product'].with_context(active_test=False).search([('attribute_value_ids', 'in', self.ids)])
        if linked_products:
            raise UserError(_('The operation cannot be completed:\nYou are trying to delete an attribute value with a reference on a product variant.'))
        return super(ProductAttributevalue, self).unlink()

    */})
pool.ProductAttributevalue().Methods().VariantName().DeclareMethod(
`VariantName` ,
func (rs pool.ProductAttributevalueSet , args struct{VariableAttributes interface{}
}){
  //@api.multi
  /*def _variant_name(self, variable_attributes):
        return ", ".join([v.name for v in self.sorted(key=lambda r: r.attribute_id.name) if v.attribute_id in variable_attributes])


*/})


pool.ProductAttributePrice().DeclareModel()
pool.ProductAttributePrice().AddMany2OneField("ProductTmplId",models.ForeignKeyFieldParams{String :"Product Template" , RelationModel: pool.ProductTemplate(), OnDelete : models.Cascade, Required: true})
pool.ProductAttributePrice().AddMany2OneField("ValueId",models.ForeignKeyFieldParams{String :"Product Attribute Value" , RelationModel: pool.ProductAttributeValue(), OnDelete : models.Cascade, Required: true})
pool.ProductAttributePrice().AddFloatField("PriceExtra", models.FloatFieldParams{String :"Price Extra"})


pool.ProductAttributeLine().DeclareModel()
pool.ProductAttributeLine().AddMany2OneField("ProductTmplId",models.ForeignKeyFieldParams{String :"Product Template" , RelationModel: pool.ProductTemplate(), OnDelete : models.Cascade, Required: true})
pool.ProductAttributeLine().AddMany2OneField("AttributeId",models.ForeignKeyFieldParams{String :"Attribute" , RelationModel: pool.ProductAttribute(), OnDelete : models.Restrict, Required: true})
pool.ProductAttributeLine().AddMany2ManyField("ValueIds", models.Many2ManyFieldParams{String :"Attribute Values" , RelationModel: pool.ProductAttributeValue()})
pool.ProductAttributeLine().Methods().CheckValidAttribute().DeclareMethod(
`CheckValidAttribute` ,
func (rs pool.ProductAttributeLineSet){
  //@api.constrains('value_ids','attribute_id')
  /*def _check_valid_attribute(self):
        if any(line.value_ids > line.attribute_id.value_ids for line in self):
            raise ValidationError(_('Error ! You cannot use this attribute with the following value.'))
        return True

    */})
pool.ProductAttributeLine().Methods().NameSearch().DeclareMethod(
`NameSearch` ,
func (rs pool.ProductAttributeLineSet , args struct{Name interface{}
Args interface{}
Operator interface{}
Limit interface{}
}){
  //@api.model
  /*def name_search(self, name='', args=None, operator='ilike', limit=100):
        # TDE FIXME: currently overriding the domain; however as it includes a
        # search on a m2o and one on a m2m, probably this will quickly become
        # difficult to compute - check if performance optimization is required
        if name and operator in ('=', 'ilike', '=ilike', 'like', '=like'):
            new_args = ['|', ('attribute_id', operator, name), ('value_ids', operator, name)]
        else:
            new_args = args
        return super(ProductAttributeLine, self).name_search(name=name, args=new_args, operator=operator, limit=limit)
*/})
 
 }