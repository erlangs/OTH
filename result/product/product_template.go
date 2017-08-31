package product 

  import (
"github.com/hexya-erp/hexya/hexya/models/types"

) 

 func init() { 

 

pool.ProductTemplate().DeclareModel()
pool.ProductTemplate().Methods().GetDefaultCategoryId().DeclareMethod(
`GetDefaultCategoryId` ,
func (rs pool.ProductTemplateSet){
  /*def _get_default_category_id(self):
        if self._context.get('categ_id') or self._context.get('default_categ_id'):
            return self._context.get('categ_id') or self._context.get('default_categ_id')
        category = self.env.ref('product.product_category_all', raise_if_not_found=False)
        return category and category.id or False

    */})
pool.ProductTemplate().Methods().GetDefaultUomId().DeclareMethod(
`GetDefaultUomId` ,
func (rs pool.ProductTemplateSet){
  /*def _get_default_uom_id(self):
        return self.env["product.uom"].search([], limit=1, order='id').id

    name = */})
pool.ProductTemplate().AddCharField("Name", models.StringFieldParams{String :"Name", Index: true, Required : true, Translate: true})
pool.ProductTemplate().AddIntegerField("Sequence", models.SimpleFieldParams{String :"Sequence", Default: func(models.Environment, models.FieldMap) interface{} {return 1} ,Help :"'Gives the sequence order when displaying a product list'"})
pool.ProductTemplate().AddTextField("Description" , models.StringFieldParams{String :"Description", Translate: true ,Help :"A precise description of the Product, used only for internal information purposes."})
pool.ProductTemplate().AddTextField("DescriptionPurchase" , models.StringFieldParams{String :"Purchase Description", Translate: true ,Help :"A description of the Product that you want to communicate to your vendors.  This description will be copied to every Purchase Order, Receipt and Vendor Bill/Refund."})
pool.ProductTemplate().AddTextField("DescriptionSale" , models.StringFieldParams{String :"Sale Description", Translate: true ,Help :"A description of the Product that you want to communicate to your customers.  This description will be copied to every Sale Order, Delivery Order and Customer Invoice/Refund"})
pool.ProductTemplate().AddSelectionField("Type", models.SelectionFieldParams{String :"Type", Selection : types.Selection{
"consu" : "_('Consumable'",
"service" : "_('Service'",
}, Default: func(models.Environment, models.FieldMap) interface{} {return "consu"}, Required : true, Help : "A stockable product is a product for which you manage stock. The "Inventory" app has to be installed.\n' 'A consumable product"})
pool.ProductTemplate().AddBooleanField("Rental", models.SimpleFieldParams{String :"Can be Rent')"})
pool.ProductTemplate().AddMany2OneField("Categ",models.ForeignKeyFieldParams{String :"Internal Category" , RelationModel: pool.ProductCategory() , JSON : "categ_id", Default : func(models.Environment, models.FieldMap) interface{}{
/*_get_default_category_id(self):
        if self._context.get('categ_id') or self._context.get('default_categ_id'):
            return self._context.get('categ_id') or self._context.get('default_categ_id')
        category = self.env.ref('product.product_category_all', raise_if_not_found=False)
        return category and category.id or False

    */
return 0}})
pool.ProductTemplate().AddMany2OneField("Currency",models.ForeignKeyFieldParams{String :"Currency" , RelationModel: pool.Currency() , JSON : "currency_id"})
pool.ProductTemplate().AddFloatField("Price", models.FloatFieldParams{String :"Price"})
pool.ProductTemplate().AddFloatField("ListPrice", models.FloatFieldParams{String :"Sale Price", Default: func(models.Environment, models.FieldMap) interface{} {return 1.0}})
pool.ProductTemplate().AddFloatField("LstPrice", models.FloatFieldParams{String :"Public Price"})
pool.ProductTemplate().AddFloatField("StandardPrice", models.FloatFieldParams{String :"Cost"})
pool.ProductTemplate().AddFloatField("Volume", models.FloatFieldParams{String :"Volume" ,Help :"The volume in m3., store"})
pool.ProductTemplate().AddFloatField("Weight", models.FloatFieldParams{String :"Weight"})
pool.ProductTemplate().AddFloatField("Warranty", models.FloatFieldParams{String :"Warranty')"})
pool.ProductTemplate().AddBooleanField("SaleOk", models.SimpleFieldParams{String :"Can be Sold", Default: func(models.Environment, models.FieldMap) interface{} {return true} ,Help :"Specify if the product can be selected in a sales order line."})
pool.ProductTemplate().AddBooleanField("PurchaseOk", models.SimpleFieldParams{String :"Can be Purchased", Default: func(models.Environment, models.FieldMap) interface{} {return true}})
pool.ProductTemplate().AddMany2OneField("Pricelist",models.ForeignKeyFieldParams{String :"Pricelist" , RelationModel: pool.ProductPricelist() , JSON : "pricelist_id" , Help :"'Technical field. Used for searching on pricelists, not stored in database.'"})
pool.ProductTemplate().AddMany2OneField("Uom",models.ForeignKeyFieldParams{String :"Unit of Measure" , RelationModel: pool.ProductUom() , JSON : "uom_id", Default : func(models.Environment, models.FieldMap) interface{}{
/*_get_default_uom_id(self):
        return self.env["product.uom"].search([], limit=1, order='id').id

    name = */
return 0}, Required : true , Help :"Default Unit of Measure used for all stock operation."})
pool.ProductTemplate().AddMany2OneField("UomPo",models.ForeignKeyFieldParams{String :"Purchase Unit of Measure" , RelationModel: pool.ProductUom() , JSON : "uom_po_id", Default : func(models.Environment, models.FieldMap) interface{}{
/*_get_default_uom_id(self):
        return self.env["product.uom"].search([], limit=1, order='id').id

    name = */
return 0}, Required : true , Help :"Default Unit of Measure used for purchase orders. It must be in the same category than the default unit of measure."})
pool.ProductTemplate().AddMany2OneField("Company",models.ForeignKeyFieldParams{String :"Company" , RelationModel: pool.Company() , JSON : "company_id", Default : func(models.Environment, models.FieldMap) interface{}{
/*lambda self: self.env['res.company']._company_default_get*/
return 0}})
pool.ProductTemplate().AddOne2ManyField("Packagings", models.ReverseFieldParams{String :"Logistical Units" ,RelationModel : pool.ProductPackaging() ,ReverseFK : "ProductTmpl" , JSON : "packaging_ids"})
pool.ProductTemplate().AddOne2ManyField("Sellers", models.ReverseFieldParams{String :"Vendors" ,RelationModel : pool.ProductSupplierinfo() ,ReverseFK : "ProductTmpl" , JSON : "seller_ids"})
pool.ProductTemplate().AddBooleanField("Active", models.SimpleFieldParams{String :"Active", Default: func(models.Environment, models.FieldMap) interface{} {return true} ,Help :"If unchecked, it will allow you to hide the product without removing it."})
pool.ProductTemplate().AddIntegerField("Color", models.SimpleFieldParams{String :"Color Index')"})
pool.ProductTemplate().AddOne2ManyField("AttributeLines", models.ReverseFieldParams{String :"Product Attributes" ,RelationModel : pool.ProductAttributeLine() ,ReverseFK : "ProductTmpl" , JSON : "attribute_line_ids"})
pool.ProductTemplate().AddOne2ManyField("ProductVariants", models.ReverseFieldParams{String :"Products" ,RelationModel : pool.ProductProduct() ,ReverseFK : "ProductTmpl" , JSON : "product_variant_ids"})
pool.ProductTemplate().AddMany2OneField("ProductVariant",models.ForeignKeyFieldParams{String :"Product" , RelationModel: pool.ProductProduct() , JSON : "product_variant_id"})
pool.ProductTemplate().AddIntegerField("ProductVariantCount", models.SimpleFieldParams{String :"# Product Variants"})
pool.ProductTemplate().AddCharField("Barcode", models.StringFieldParams{String :"Barcode"})
pool.ProductTemplate().AddCharField("DefaultCode", models.StringFieldParams{String :"Internal Reference"})
pool.ProductTemplate().AddOne2ManyField("Items", models.ReverseFieldParams{String :"Pricelist Items" ,RelationModel : pool.ProductPricelistItem() ,ReverseFK : "ProductTmpl" , JSON : "item_ids"})
pool.ProductTemplate().AddBinaryField("Image", models.SimpleFieldParams{String :"Image" ,Help :"This field holds the image used as image for the product, limited to 1024x1024px."})
pool.ProductTemplate().AddBinaryField("ImageMedium", models.SimpleFieldParams{String :"Medium-sized image" ,Help :"Medium-sized image of the product. It is automatically  resized as a 128x128px image, with aspect ratio preserved,  only when the image exceeds one of those sizes. Use this field in form views or some kanban views."})
pool.ProductTemplate().AddBinaryField("ImageSmall", models.SimpleFieldParams{String :"Small-sized image" ,Help :"Small-sized image of the product. It is automatically  resized as a 64x64px image, with aspect ratio preserved.  Use this field anywhere a small image is required."})
pool.ProductTemplate().Methods().ComputeProductVariantId().DeclareMethod(
`ComputeProductVariantId` ,
func (rs pool.ProductTemplateSet){
  //@api.depends('product_variant_ids')
  /*def _compute_product_variant_id(self):
        for p in self:
            p.product_variant_id = p.product_variant_ids[:1].id

    */})
pool.ProductTemplate().Methods().ComputeCurrencyId().DeclareMethod(
`ComputeCurrencyId` ,
func (rs pool.ProductTemplateSet){
  //@api.multi
  /*def _compute_currency_id(self):
        try:
            main_company = self.sudo().env.ref('base.main_company')
        except ValueError:
            main_company = self.env['res.company'].sudo().search([], limit=1, order="id")
        for template in self:
            template.currency_id = template.company_id.sudo().currency_id.id or main_company.currency_id.id

    */})
pool.ProductTemplate().Methods().ComputeTemplatePrice().DeclareMethod(
`ComputeTemplatePrice` ,
func (rs pool.ProductTemplateSet){
  //@api.multi
  /*def _compute_template_price(self):
        prices = {}
        pricelist_id_or_name = self._context.get('pricelist')
        if pricelist_id_or_name:
            pricelist = None
            partner = self._context.get('partner')
            quantity = self._context.get('quantity', 1.0)

            # Support context pricelists specified as display_name or ID for compatibility
            if isinstance(pricelist_id_or_name, basestring):
                pricelist = self.env['product.pricelist'].name_search(pricelist_id_or_name, operator='=', limit=1)
            elif isinstance(pricelist_id_or_name, (int, long)):
                pricelist = self.env['product.pricelist'].browse(pricelist_id_or_name)

            if pricelist:
                quantities = [quantity] * len(self)
                partners = [partner] * len(self)
                prices = pricelist.get_products_price(self, quantities, partners)

        for template in self:
            template.price = prices.get(template.id, 0.0)

    */})
pool.ProductTemplate().Methods().InverseTemplatePrice().DeclareMethod(
`InverseTemplatePrice` ,
func (rs pool.ProductTemplateSet){
  //@api.multi
  /*def _set_template_price(self):
        if self._context.get('uom'):
            for template in self:
                value = self.env['product.uom'].browse(self._context['uom'])._compute_price(template.price, template.uom_id)
                template.write({'list_price': value})
        else:
            self.write({'list_price': self.price})

    */})
pool.ProductTemplate().Methods().ComputeStandardPrice().DeclareMethod(
`ComputeStandardPrice` ,
func (rs pool.ProductTemplateSet){
  //@api.depends('product_variant_ids','product_variant_ids.standard_price')
  /*def _compute_standard_price(self):
        unique_variants = self.filtered(lambda template: len(template.product_variant_ids) == 1)
        for template in unique_variants:
            template.standard_price = template.product_variant_ids.standard_price
        for template in (self - unique_variants):
            template.standard_price = 0.0

    */})
pool.ProductTemplate().Methods().InverseStandardPrice().DeclareMethod(
`InverseStandardPrice` ,
func (rs pool.ProductTemplateSet){
  //@api.one
  /*def _set_standard_price(self):
        if len(self.product_variant_ids) == 1:
            self.product_variant_ids.standard_price = self.standard_price

    */})
pool.ProductTemplate().Methods().SearchStandardPrice().DeclareMethod(
`SearchStandardPrice` ,
func (rs pool.ProductTemplateSet , args struct{Operator interface{}
Value interface{}
}){
  /*def _search_standard_price(self, operator, value):
        products = self.env['product.product'].search([('standard_price', operator, value)], limit=None)
        return [('id', 'in', products.mapped('product_tmpl_id').ids)]

    */})
pool.ProductTemplate().Methods().ComputeVolume().DeclareMethod(
`ComputeVolume` ,
func (rs pool.ProductTemplateSet){
  //@api.depends('product_variant_ids','product_variant_ids.volume')
  /*def _compute_volume(self):
        unique_variants = self.filtered(lambda template: len(template.product_variant_ids) == 1)
        for template in unique_variants:
            template.volume = template.product_variant_ids.volume
        for template in (self - unique_variants):
            template.volume = 0.0

    */})
pool.ProductTemplate().Methods().InverseVolume().DeclareMethod(
`InverseVolume` ,
func (rs pool.ProductTemplateSet){
  //@api.one
  /*def _set_volume(self):
        if len(self.product_variant_ids) == 1:
            self.product_variant_ids.volume = self.volume

    */})
pool.ProductTemplate().Methods().ComputeWeight().DeclareMethod(
`ComputeWeight` ,
func (rs pool.ProductTemplateSet){
  //@api.depends('product_variant_ids','product_variant_ids.weight')
  /*def _compute_weight(self):
        unique_variants = self.filtered(lambda template: len(template.product_variant_ids) == 1)
        for template in unique_variants:
            template.weight = template.product_variant_ids.weight
        for template in (self - unique_variants):
            template.weight = 0.0

    */})
pool.ProductTemplate().Methods().InverseWeight().DeclareMethod(
`InverseWeight` ,
func (rs pool.ProductTemplateSet){
  //@api.one
  /*def _set_weight(self):
        if len(self.product_variant_ids) == 1:
            self.product_variant_ids.weight = self.weight

    */})
pool.ProductTemplate().Methods().ComputeProductVariantCount().DeclareMethod(
`ComputeProductVariantCount` ,
func (rs pool.ProductTemplateSet){
  //@api.depends('product_variant_ids.product_tmpl_id')
  /*def _compute_product_variant_count(self):
        self.product_variant_count = len(self.product_variant_ids)

    */})
pool.ProductTemplate().Methods().ComputeDefaultCode().DeclareMethod(
`ComputeDefaultCode` ,
func (rs pool.ProductTemplateSet){
  //@api.depends('product_variant_ids','product_variant_ids.default_code')
  /*def _compute_default_code(self):
        unique_variants = self.filtered(lambda template: len(template.product_variant_ids) == 1)
        for template in unique_variants:
            template.default_code = template.product_variant_ids.default_code
        for template in (self - unique_variants):
            template.default_code = ''

    */})
pool.ProductTemplate().Methods().InverseDefaultCode().DeclareMethod(
`InverseDefaultCode` ,
func (rs pool.ProductTemplateSet){
  //@api.one
  /*def _set_default_code(self):
        if len(self.product_variant_ids) == 1:
            self.product_variant_ids.default_code = self.default_code

    */})
pool.ProductTemplate().Methods().CheckUom().DeclareMethod(
`CheckUom` ,
func (rs pool.ProductTemplateSet){
  //@api.constrains('uom_id','uom_po_id')
  /*def _check_uom(self):
        if any(template.uom_id and template.uom_po_id and template.uom_id.category_id != template.uom_po_id.category_id for template in self):
            raise ValidationError(_('Error: The default Unit of Measure and the purchase Unit of Measure must be in the same category.'))
        return True

    */})
pool.ProductTemplate().Methods().OnchangeUomId().DeclareMethod(
`OnchangeUomId` ,
func (rs pool.ProductTemplateSet){
  //@api.onchange('uom_id')
  /*def _onchange_uom_id(self):
        if self.uom_id:
            self.uom_po_id = self.uom_id.id

    */})
pool.ProductTemplate().Methods().Create().DeclareMethod(
`Create` ,
func (rs pool.ProductTemplateSet , args struct{}){
  //@api.model
  /*def create(self, vals):
        ''' Store the initial standard price in order to be able to retrieve the cost of a product template for a given date'''
        # TDE FIXME: context brol
        tools.image_resize_images(vals)
        template = super(ProductTemplate, self).create(vals)
        if "create_product_product" not in self._context:
            template.create_variant_ids()

        # This is needed to set given values to first variant after creation
        related_vals = {}
        if vals.get('barcode'):
            related_vals['barcode'] = vals['barcode']
        if vals.get('default_code'):
            related_vals['default_code'] = vals['default_code']
        if vals.get('standard_price'):
            related_vals['standard_price'] = vals['standard_price']
        if vals.get('volume'):
            related_vals['volume'] = vals['volume']
        if vals.get('weight'):
            related_vals['weight'] = vals['weight']
        if related_vals:
            template.write(related_vals)
        return template

    */})
pool.ProductTemplate().Methods().Write().DeclareMethod(
`Write` ,
func (rs pool.ProductTemplateSet , args struct{}){
  //@api.multi
  /*def write(self, vals):
        tools.image_resize_images(vals)
        res = super(ProductTemplate, self).write(vals)
        if 'attribute_line_ids' in vals or vals.get('active'):
            self.create_variant_ids()
        if 'active' in vals and not vals.get('active'):
            self.with_context(active_test=False).mapped('product_variant_ids').write({'active': vals.get('active')})
        return res

    */})
pool.ProductTemplate().Methods().Copy().DeclareMethod(
`Copy` ,
func (rs pool.ProductTemplateSet , args struct{Default interface{}
}){
  //@api.multi
  /*def copy(self, default=None):
        # TDE FIXME: should probably be copy_data
        self.ensure_one()
        if default is None:
            default = {}
        if 'name' not in default:
            default['name'] = _("%s (copy)") % self.name
        return super(ProductTemplate, self).copy(default=default)

    */})
pool.ProductTemplate().Methods().NameGet().DeclareMethod(
`NameGet` ,
func (rs pool.ProductTemplateSet){
  //@api.multi
  /*def name_get(self):
        return [(template.id, '%s%s' % (template.default_code and '[%s] ' % template.default_code or '', template.name))
                for template in self]

    */})
pool.ProductTemplate().Methods().NameSearch().DeclareMethod(
`NameSearch` ,
func (rs pool.ProductTemplateSet , args struct{Name interface{}
Args interface{}
Operator interface{}
Limit interface{}
}){
  //@api.model
  /*def name_search(self, name='', args=None, operator='ilike', limit=100):
        # Only use the product.product heuristics if there is a search term and the domain
        # does not specify a match on `product.template` IDs.
        if not name or any(term[0] == 'id' for term in (args or [])):
            return super(ProductTemplate, self).name_search(name=name, args=args, operator=operator, limit=limit)

        Product = self.env['product.product']
        templates = self.browse([])
        while True:
            domain = templates and [('product_tmpl_id', 'not in', templates.ids)] or []
            args = args if args is not None else []
            products_ns = Product.name_search(name, args+domain, operator=operator)
            products = Product.browse([x[0] for x in products_ns])
            templates |= products.mapped('product_tmpl_id')
            if (not products) or (limit and (len(templates) > limit)):
                break

        # re-apply product.template order + name_get
        return super(ProductTemplate, self).name_search(
            '', args=[('id', 'in', list(set(templates.ids)))],
            operator='ilike', limit=limit)

    */})
pool.ProductTemplate().Methods().PriceCompute().DeclareMethod(
`PriceCompute` ,
func (rs pool.ProductTemplateSet , args struct{PriceType interface{}
Uom interface{}
Currency interface{}
Company interface{}
}){
  //@api.multi
  /*def price_compute(self, price_type, uom=False, currency=False, company=False):
        # TDE FIXME: delegate to template or not ? fields are reencoded here ...
        # compatibility about context keys used a bit everywhere in the code
        if not uom and self._context.get('uom'):
            uom = self.env['product.uom'].browse(self._context['uom'])
        if not currency and self._context.get('currency'):
            currency = self.env['res.currency'].browse(self._context['currency'])

        templates = self
        if price_type == 'standard_price':
            # standard_price field can only be seen by users in base.group_user
            # Thus, in order to compute the sale price from the cost for users not in this group
            # We fetch the standard price as the superuser
            templates = self.with_context(force_company=company and company.id or self._context.get('force_company', self.env.user.company_id.id)).sudo()

        prices = dict.fromkeys(self.ids, 0.0)
        for template in templates:
            prices[template.id] = template[price_type] or 0.0

            if uom:
                prices[template.id] = template.uom_id._compute_price(prices[template.id], uom)

            # Convert from current user company currency to asked one
            # This is right cause a field cannot be in more than one currency
            if currency:
                prices[template.id] = template.currency_id.compute(prices[template.id], currency)

        return prices

    # compatibility to remove after v10 - DEPRECATED
    */})
pool.ProductTemplate().Methods().PriceGet().DeclareMethod(
`PriceGet` ,
func (rs pool.ProductTemplateSet , args struct{Products interface{}
Ptype interface{}
}){
  //@api.model
  /*def _price_get(self, products, ptype='list_price'):
        return products.price_compute(ptype)

    */})
pool.ProductTemplate().Methods().CreateVariantIds().DeclareMethod(
`CreateVariantIds` ,
func (rs pool.ProductTemplateSet){
  //@api.multi
  /*def create_variant_ids(self):
        Product = self.env["product.product"]
        for tmpl_id in self.with_context(active_test=False):
            # adding an attribute with only one value should not recreate product
            # write this attribute on every product to make sure we don't lose them
            variant_alone = tmpl_id.attribute_line_ids.filtered(lambda line: len(line.value_ids) == 1).mapped('value_ids')
            for value_id in variant_alone:
                updated_products = tmpl_id.product_variant_ids.filtered(lambda product: value_id.attribute_id not in product.mapped('attribute_value_ids.attribute_id'))
                updated_products.write({'attribute_value_ids': [(4, value_id.id)]})

            # list of values combination
            existing_variants = [set(variant.attribute_value_ids.ids) for variant in tmpl_id.product_variant_ids]
            variant_matrix = itertools.product(*(line.value_ids for line in tmpl_id.attribute_line_ids if line.value_ids and line.value_ids[0].attribute_id.create_variant))
            variant_matrix = map(lambda record_list: reduce(lambda x, y: x+y, record_list, self.env['product.attribute.value']), variant_matrix)
            to_create_variants = filter(lambda rec_set: set(rec_set.ids) not in existing_variants, variant_matrix)

            # check product
            variants_to_activate = self.env['product.product']
            variants_to_unlink = self.env['product.product']
            for product_id in tmpl_id.product_variant_ids:
                if not product_id.active and product_id.attribute_value_ids in variant_matrix:
                    variants_to_activate |= product_id
                elif product_id.attribute_value_ids not in variant_matrix:
                    variants_to_unlink |= product_id
            if variants_to_activate:
                variants_to_activate.write({'active': True})

            # create new product
            for variant_ids in to_create_variants:
                new_variant = Product.create({
                    'product_tmpl_id': tmpl_id.id,
                    'attribute_value_ids': [(6, 0, variant_ids.ids)]
                })

            # unlink or inactive product
            for variant in variants_to_unlink:
                try:
                    with self._cr.savepoint(), tools.mute_logger('odoo.sql_db'):
                        variant.unlink()
                # We catch all kind of exception to be sure that the operation doesn't fail.
                except (psycopg2.Error, except_orm):
                    variant.write({'active': False})
                    pass
        return True
*/})
 
 }