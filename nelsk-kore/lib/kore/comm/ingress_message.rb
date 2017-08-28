module Kore
  module Comm
    class IngressMessage
      attr_accessor :originator, :plugin, :cmd, :content, :raw
      def initialize(o)
        @originator = o.fetch(:originator, '')
        @plugin = o.fetch(:plugin, '')
        @content = o.fetch(:content, '')
      end

      def to_s
        "[#{@originator}] [#{@plugin}] [#{@content}]"
      end
    end
  end
end
